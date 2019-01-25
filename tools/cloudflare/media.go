package cloudflare

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/xid"
)

// RegisterMediaWebhook registers a url as media webhook
// Reference: https://developers.cloudflare.com/stream/webhooks/
func (c *Client) RegisterMediaWebhook(baseURL string) (relativePath, absoluteURL string, err error) {
	// Request model
	type Request struct {
		NotificationURL string `json:"notification_url"`
	}

	// Response model
	type Response struct {
		Result struct {
			NotificationURL string    `json:"notification_url"`
			ModifiedAt      time.Time `json:"modified"`
		} `json:"result"`
		Success  bool           `json:"success"`
		Errors   []GeneralError `json:"errors"`
		Messages []string       `json:"messages"`
	}

	// Generate uinque id
	uid := xid.New().String()
	webhookPath := "/cloudflare/media/" + uid
	webhookURL := baseURL + webhookPath

	// Media Webhook Endpoint
	endpoint := c.BaseURL + "/accounts/" + c.AccountID + "/media/webhook"

	// New Request
	req, err := json.Marshal(Request{
		NotificationURL: webhookURL,
	})
	if err != nil {
		return "", "", err
	}
	reqBuffer := bytes.NewBuffer(req)

	// Create http request
	httpReq, err := http.NewRequest(http.MethodPut, endpoint, reqBuffer)
	if err != nil {
		return "", "", err
	}
	httpReq.Header.Add("X-Auth-Key", c.APIKey)
	httpReq.Header.Add("X-Auth-Email", c.Email)

	// Perform the request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", "", err
	}

	// Read response body
	respData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return "", "", err
	}

	// Unmarshal response
	var resp Response
	if err := json.Unmarshal(respData, &resp); err != nil {
		return "", "", err
	}

	// Checking if API call was successful
	if !resp.Success {
		return "", "", resp.Errors[0]
	}

	return webhookPath, resp.Result.NotificationURL, nil
}
