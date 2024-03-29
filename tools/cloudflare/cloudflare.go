package cloudflare

import (
	"fmt"
	"net/http"
	"os"
)

// A Credentials contains required authentication credentials
type Credentials struct {
	Email, APIKey string
}

// A Client contains all required values for services
type Client struct {
	BaseURL   string
	AccountID string
	Credentials
	httpClient *http.Client
}

// DefaultClient is the default CloudFlare client
var DefaultClient = NewClient(
	os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
	os.Getenv("CLOUDFLARE_EMAIL"),
	os.Getenv("CLOUDFLARE_API_KEY"),
)

// An GeneralError cotains a code and a message
type GeneralError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error method implements error interface
func (err GeneralError) Error() string {
	return fmt.Sprintf("%v (%v)", err.Message, err.Code)
}

// NewClient creates new CloudFlare client
func NewClient(accountID, email, apiKey string) *Client {
	return &Client{
		BaseURL:   "https://api.cloudflare.com/client/v4",
		AccountID: accountID,
		Credentials: Credentials{
			Email:  email,
			APIKey: apiKey,
		},
		httpClient: http.DefaultClient,
	}
}
