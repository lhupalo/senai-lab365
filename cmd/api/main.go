package main

import (
	"senai-lab365/internal/application"
	"senai-lab365/internal/infrastructure"
	"senai-lab365/internal/interfaces/handlers"

	_ "senai-lab365/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Notifications API
// @version         1.0
// @description     API de Notificações Assíncronas
// @host            localhost:8000
// @BasePath        /v1
func main() {
	dispatcher := infrastructure.NewNotificationDispatcher(5, 100)
	defer dispatcher.Shutdown()

	uc := application.NewSendNotificationUseCase(dispatcher)
	handler := handlers.NewNotificationHandler(uc)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/notifications", handler.Create)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	router.Run(":8000")
}
