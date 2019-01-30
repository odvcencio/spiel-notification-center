package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"spiel/notification-center/database"
	"spiel/notification-center/tools/onesignal"
	"time"

	"github.com/labstack/echo"
	nsq "github.com/nsqio/go-nsq"
)

func handleMuxMediaNotification(ctx echo.Context) error {
	type Request struct {
		Type   string `json:"type"`
		ID     string `json:"id"`
		Object struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"object"`
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

//
// func handleCloudFlareMediaNotification(ctx echo.Context) error {
// 	type Request struct {
// 		UID            string `json:"uid"`
// 		ThumbnailImage string `json:"thumbnail"`
// 		ReadyToStream  bool   `json:"readyToStream"`
// 		Status         struct {
// 			State string `json:"state"`
// 		} `json:"status"`
// 		Meta       map[string]string `json:"meta"`
// 		Labels     []string          `json:"labels"`
// 		CreatedAt  time.Time         `json:"created"`
// 		ModifiedAt time.Time         `json:"modified"`
// 		Size       uint32            `json:"size"`
// 		PreviewURL string            `json:"preview"`
// 	}
//
// 	// Read request body
// 	reqData, err := ioutil.ReadAll(ctx.Request().Body)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
//
// 	// Unmarshal request json
// 	var req Request
// 	if err := json.Unmarshal(reqData, &req); err != nil {
// 		println(err)
// 		return err
// 	}
//
// 	// TODO: Create appropriate Spiel and connect it to
// 	//       appropriate User and Question
//
// 	// Responsing with OK
// 	ctx.NoContent(200)
//
// 	return nil
// }

func handleTopicQuestionToUser(message *nsq.Message) error {
	// Message model
	var msg struct {
		QuestionID int    `json:"question_id"`
		UserID     string `json:"user_id"`
	}

	// Decoding data
	if err := json.Unmarshal(message.Body, &msg); err != nil {
		return err
	}

	// Find user info
	user, err := database.GetUserByID(msg.UserID)
	if err != nil {
		return err
	}

	// Find question info
	question, err := database.GetQuestionByID(msg.QuestionID)
	if err != nil {
		return err
	}

	// Sending notification
	onesignal.DefaultClient.SendPushNotification(onesignal.Notification{
		Contents: map[string]string{
			"en": question.Question,
		},
		Headings: map[string]string{
			"en": question.User.FirstName + " " +
				question.User.LastName + " asked you a question",
		},
		Filters: []interface{}{
			onesignal.Filter{
				Field:    "tag",
				Key:      "user_id",
				Relation: "=",
				Value:    user.ID,
			},
		},
	})

	// Print response
	log.Println(string(message.Body))

	return nil
}
