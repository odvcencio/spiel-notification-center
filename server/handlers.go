package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/labstack/echo"
)

func handleCloudFlareMediaNotification(ctx echo.Context) error {
	type Request struct {
		UID            string `json:"uid"`
		ThumbnailImage string `json:"thumbnail"`
		ReadyToStream  bool   `json:"readyToStream"`
		Status         struct {
			State string `json:"state"`
		} `json:"status"`
		Meta       map[string]string `json:"meta"`
		Labels     []string          `json:"labels"`
		CreatedAt  time.Time         `json:"created"`
		ModifiedAt time.Time         `json:"modified"`
		Size       uint32            `json:"size"`
		PreviewURL string            `json:"preview"`
	}

	// Read request body
	reqData, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return err
	}

	// Unmarshal request json
	var req Request
	if err := json.Unmarshal(reqData, &req); err != nil {
		println(err)
		return err
	}

	// TODO: Create appropriate Spiel and connect it to
	//       appropriate User and Question

	// Responsing with OK
	ctx.NoContent(200)

	return nil
}
