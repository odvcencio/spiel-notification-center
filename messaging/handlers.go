package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"spiel/notification-center/database"
	"spiel/notification-center/models"
	"spiel/notification-center/tools/onesignal"
	tools "spiel/notification-center/tools/sendgrid"
	"strings"
	"time"

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
		AssessmentID int    `json:"assessment_id"`
		UserID       string `json:"user_id"`
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

	var assessmentNotification models.Notification

	assessmentNotification.SpielID = assessment.SpielID
	assessmentNotification.UserID = msg.UserID
	assessmentNotification.SpielAssessmentID = msg.AssessmentID
	assessmentNotification.Message = assessment.User.FirstName + " has given you a Spiel assessment!"
	assessmentNotification.Type = "assessment"
	assessmentNotification.CreatedTime = time.Now()

	if err := database.InsertNotificationForSpiel(assessmentNotification); err != nil {
		log.Println(err)
	}

	// Sending notification
	onesignal.DefaultClient.SendPushNotification(onesignal.Notification{
		Data: map[string]string{
			"type": assessmentNotification.Type,
		},
		Contents: map[string]string{
			"en": assessment.User.FirstName + " has given you a Spiel assessment!",
		},
		Headings: map[string]string{
			"en": "Someone has assessed your Spiel!",
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
