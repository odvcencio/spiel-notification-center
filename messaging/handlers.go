package messaging

import (
	"encoding/json"
	"log"
	"spiel/notification-center/database"
	"spiel/notification-center/tools/onesignal"

	nsq "github.com/nsqio/go-nsq"
)

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