package application

import (
	"time"

	"senai-lab365/internal/domain"

	"github.com/google/uuid"
)

type SendNotificationInput struct {
	UserID   string
	Message  string
	Priority string
}

type SendNotificationUseCase struct {
	queue domain.NotificationQueue
}

func NewSendNotificationUseCase(queue domain.NotificationQueue) *SendNotificationUseCase {
	return &SendNotificationUseCase{queue: queue}
}

func (uc *SendNotificationUseCase) Execute(input SendNotificationInput) (*domain.Notification, error) {
	priority := domain.Priority(input.Priority)
	if priority != domain.PriorityLow && priority != domain.PriorityMedium && priority != domain.PriorityHigh {
		priority = domain.PriorityMedium
	}

	notification := &domain.Notification{
		ID:        uuid.New().String(),
		UserID:    input.UserID,
		Message:   input.Message,
		Priority:  priority,
		CreatedAt: time.Now(),
	}

	if err := uc.queue.Enqueue(notification); err != nil {
		return nil, err
	}

	return notification, nil
}
