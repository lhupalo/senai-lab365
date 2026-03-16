package infrastructure

import (
	"context"
	"log"
	"senai-lab365/internal/domain"
	"sync"
	"time"
)

type NotificationDispatcher struct {
	workerCount  int
	queue        chan *domain.Notification
	metricsChan  chan int
	ctx          context.Context
	cancel       context.CancelFunc
	wg           sync.WaitGroup
}

func NewNotificationDispatcher(workerCount, bufferSize int) *NotificationDispatcher {
	ctx, cancel := context.WithCancel(context.Background())
	d := &NotificationDispatcher{
		workerCount: workerCount,
		queue:       make(chan *domain.Notification, bufferSize),
		metricsChan:  make(chan int),
		ctx:         ctx,
		cancel:      cancel,
	}
	d.startWorkerPool()
	return d
}

func (d *NotificationDispatcher) startWorkerPool() {
	for i := 0; i < d.workerCount; i++ {
		d.wg.Add(1)
		go d.worker(i)
	}
}

func (d *NotificationDispatcher) worker(id int) {
	defer d.wg.Done()
	for {
		select {
		case <-d.ctx.Done():
			return
		case notification, ok := <-d.queue:
			if !ok {
				return
			}
			d.dispatch(notification, id)
		}
	}
}

func (d *NotificationDispatcher) dispatch(notification *domain.Notification, workerID int) {
	go d.recordMetric(workerID)
	time.Sleep(50 * time.Millisecond)
	log.Printf("[worker-%d] dispatched: user=%s priority=%s msg=%s",
		workerID, notification.UserID, notification.Priority, notification.Message)
}

func (d *NotificationDispatcher) recordMetric(workerID int) {
	d.metricsChan <- workerID
}

func (d *NotificationDispatcher) Enqueue(notification *domain.Notification) error {
	select {
	case <-d.ctx.Done():
		return d.ctx.Err()
	case d.queue <- notification:
		return nil
	}
}

func (d *NotificationDispatcher) Shutdown() {
	d.cancel()
	close(d.queue)
	d.wg.Wait()
}
