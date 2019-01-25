package server

import (
	"log"
	"os"
	"spiel/notification-center/tools/cloudflare"

	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	// Webhook Group
	webhook := e.Group("/webhook")

	// Registering CloudFlare Media Webhook
	log.Println("Registering CloudFlare Media Webhook...")
	cfClient := cloudflare.NewClient(
		os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		os.Getenv("CLOUDFLARE_EMAIL"),
		os.Getenv("CLOUDFLARE_API_KEY"),
	)
	externalURL := os.Getenv("NOTIFICATION_CENTER_EXTERNAL_URL")
	if relativePath, absoluteURL, err := cfClient.RegisterMediaWebhook(externalURL + "/webhook"); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("CloudFlare Media Webhook registered successfully. (%v)\n", absoluteURL)
		webhook.POST(relativePath, handleCloudFlareMediaNotification)
	}
}
