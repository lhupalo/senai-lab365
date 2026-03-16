package dto

type CreateNotificationRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Message  string `json:"message" binding:"required"`
	Priority string `json:"priority" binding:"required,oneof=low medium high"`
}
