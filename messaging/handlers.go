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
			"type": "question",
			"id":   fmt.Sprintf("%d", question.ID),
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

func handleTopicSpielAssessment(message *nsq.Message) error {
	// Message model
	var msg struct {
		AssessmentID int `json:"assessment_id"`
		UserID       int `json:"user_id"`
	}

	// Decoding data
	if err := json.Unmarshal(message.Body, &msg); err != nil {
		return err
	}

	// Find spiel assessment
	assessment, err := database.GetSpielAssessmentByID(msg.AssessmentID)
	if err != nil {
		return err
	}

	// Sending notification
	onesignal.DefaultClient.SendPushNotification(onesignal.Notification{
		Data: map[string]string{
			"type": "assessment",
		},
		Contents: map[string]string{
			"en": "Someone has assessed your Spiel!",
		},
		Headings: map[string]string{
			"en": assessment.User.FirstName + " has given you a Spiel assessment!",
		},
		Filters: []interface{}{
			onesignal.Filter{
				Field:    "tag",
				Key:      "user_id",
				Relation: "=",
				Value:    msg.UserID,
			},
		},
	})

	// Print response
	log.Println(string(message.Body))

	return nil
}
