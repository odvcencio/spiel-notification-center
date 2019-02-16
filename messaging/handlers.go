package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"spiel/notification-center/database"
	"spiel/notification-center/tools/onesignal"
	tools "spiel/notification-center/tools/sendgrid"
	"strings"

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

	if strings.Contains(question.UserID, "@") {
		tools.SendEmailPromptToWebUser(question.User, user)
	}

	// Sending notification
	onesignal.DefaultClient.SendPushNotification(onesignal.Notification{
		Data: map[string]string{
			"type":        "question",
			"question_id": fmt.Sprintf("%d", question.ID),
		},
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
