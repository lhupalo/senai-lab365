package domain

type NotificationQueue interface {
	Enqueue(notification *Notification) error
}
