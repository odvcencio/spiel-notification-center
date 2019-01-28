package onesignal

import (
	"net/http"
	"os"
)

// A Credentials contains required authentication credentials
type Credentials struct {
	AppID, APIKey string
}

// A Client contains all required values for services
type Client struct {
	BaseURL string
	Credentials
	httpClient *http.Client
}

// DefaultClient is the default OneSignal client
var DefaultClient = NewClient(
	os.Getenv("ONESIGNAL_APP_ID"),
	os.Getenv("ONESIGNAL_API_KEY"),
)

// NewClient creates new OneSignal client
func NewClient(appID, apiKey string) *Client {
	return &Client{
		BaseURL: "https://onesignal.com/api/v1",
		Credentials: Credentials{
			AppID:  appID,
			APIKey: apiKey,
		},
		httpClient: http.DefaultClient,
	}
}
