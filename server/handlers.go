package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"spiel/notification-center/database"
	"spiel/notification-center/models"
	"spiel/notification-center/tools/onesignal"
	"strings"
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
		videoID := dict["object"].(map[string]interface{})["id"].(string)

		dataMap := dict["data"].(map[string]interface{})
		playbackIDArray := dataMap["playback_ids"].([]interface{})
		playbackIDObject := playbackIDArray[0].(map[string]interface{})
		playbackID := playbackIDObject["id"].(string)

		videoURL := fmt.Sprintf("https://stream.mux.com/%s.m3u8", playbackID)

		spiel, err := database.CheckForSpiel(videoID)
		if err != nil {
			log.Println(err)
		}

		spiel.VideoURL = videoURL
		spiel.CreatedTime = time.Now()

		if err := database.UpdateSpielWithVideoURL(spiel); err != nil {
			log.Println(err)
		}

		log.Println(spiel.Question.User)

		var spielNotification models.Notification

		spielNotification.SpielID = spiel.ID
		spielNotification.UserID = spiel.Question.UserID
		spielNotification.Message = spiel.User.FirstName + " has answered your question!"

		if err := database.InsertNotificationForSpiel(spielNotification); err != nil {
			log.Println(err)
		}

		cleanName := strings.TrimSpace(spiel.User.FirstName)

		// Sending notification
		onesignal.DefaultClient.SendPushNotification(onesignal.Notification{
			Contents: map[string]string{
				"en": cleanName + " has sent you a Spiel!",
			},
			Headings: map[string]string{
				"en": "You have received a Spiel!",
			},
			Filters: []interface{}{
				onesignal.Filter{
					Field:    "tag",
					Key:      "user_id",
					Relation: "=",
					Value:    spielNotification.UserID,
				},
			},
		})

		return nil
	}

	return nil
}
