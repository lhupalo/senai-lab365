package domain

import "time"

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Notification struct {
	ID        string
	UserID    string
	Message   string
	Priority  Priority
	CreatedAt time.Time
}
