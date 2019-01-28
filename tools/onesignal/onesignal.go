package onesignal

import "net/http"

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
