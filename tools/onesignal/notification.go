package onesignal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Notification struct {
	AppID    string            `json:"app_id"`
	Contents map[string]string `json:"contents"`
	Headings map[string]string `json:"headings"`
	Subtitle map[string]string `json:"subtitle"`
	Data     map[string]string `json:"data"`
	Filters  []interface{}     `json:"filters"`
}

type Filter struct {
	Field    string `json:"field"`
	Key      string `json:"key"`
	Relation string `json:"relation"`
	Value    string `json:"value"`
}

type FilterOperator struct {
	Operator string `json:"operator"`
}

func (c *Client) SendPushNotification(notification Notification) error {
	// Send Notification Endpoint
	endpoint := c.BaseURL + "/notifications"

	// Marshaling model
	notification.AppID = c.AppID
	reqData, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	reqBuffer := bytes.NewBuffer(reqData)

	// Create http request
	httpReq, err := http.NewRequest(http.MethodPost, endpoint, reqBuffer)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic "+c.APIKey)

	// Perform the request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	// Read response body
	respData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	log.Println(string(respData))

	return nil
}
