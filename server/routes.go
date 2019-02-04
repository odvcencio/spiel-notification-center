package server

import (
	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	// Webhook Group
	webhook := e.Group("/webhook")
	webhook.POST("/videoUploaded", handleMuxMediaNotification)
}
