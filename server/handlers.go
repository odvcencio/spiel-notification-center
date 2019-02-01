package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/labstack/echo"
)

func handleMuxMediaNotification(ctx echo.Context) error {
	type Request struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Data struct {
			PlaybackIDs []struct {
				ID string `json:"id"`
			} `json:"playback_ids"`
		} `json:"data"`
		CreatedAt time.Time `json:"created_at"`
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

	log.Println(req)

	// TODO: Create appropriate Spiel and connect it to
	//       appropriate User and Question

	//database.UpdateSpielWithVideoURL(req.ID, videoID)
	// Responsing with OK
	ctx.NoContent(200)

	return nil
}
