package handlers

import (
	"net/http"
	"senai-lab365/internal/application"
	"senai-lab365/internal/interfaces/dto"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	sendNotification *application.SendNotificationUseCase
}

func NewNotificationHandler(sendNotification *application.SendNotificationUseCase) *NotificationHandler {
	return &NotificationHandler{sendNotification: sendNotification}
}

// Create godoc
// @Summary      Criar notificação
// @Description  Enfileira uma notificação para envio assíncrono (email/SMS)
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Param        body  body  object  true  "Dados da notificação"  example({"user_id":"user-123","message":"Sua fatura vence amanhã","priority":"high"})
// @Success      202   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /notifications [post]
func (h *NotificationHandler) Create(c *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := application.SendNotificationInput{
		UserID:   req.UserID,
		Message:  req.Message,
		Priority: req.Priority,
	}
	notification, err := h.sendNotification.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id":         notification.ID,
		"user_id":    notification.UserID,
		"message":    notification.Message,
		"priority":   notification.Priority,
		"created_at": notification.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}
