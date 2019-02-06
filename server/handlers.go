package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"spiel/notification-center/database"
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
	var dict echo.Map
	if err := json.Unmarshal(reqData, &dict); err != nil {
		log.Println(err)
		return err
	}

	webhookType := dict["type"].(string)

	if webhookType == "video.asset.ready" {
		videoID := dict["object"].(echo.Map)["id"].(string)

		dataMap := dict["data"].(echo.Map)
		playbackIDArray := dataMap["playback_ids"].([]echo.Map)
		playbackIDObject := playbackIDArray[0]
		playbackID := playbackIDObject["id"].(string)

		videoURL := fmt.Sprintf("https://stream.mux.com/%s.m3u8", playbackID)

		if err := database.UpdateSpielWithVideoURL(videoID, videoURL); err != nil {
			log.Println(err)
		}

		//send notification
	}

	return nil
}
