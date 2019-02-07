package database

import (
	"spiel/notification-center/models"
	"time"
)

func GetUserByID(userID string) (models.User, error) {
	var user models.User

	db = connectToDB()
	if err := db.Model(&user).
		Where("\"user\".id = ?", userID).
		Limit(1).
		Select(); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetQuestionByID(questionID int) (models.Question, error) {
	var question models.Question

	db = connectToDB()
	if err := db.Model(&question).
		Where("\"question\".id = ?", questionID).
		Relation("User").
		Limit(1).
		Select(); err != nil {
		return models.Question{}, err
	}

	return question, nil
}

func CheckForSpiel(videoID string) (models.Spiel, error) {
	var spiel models.Spiel

	db = connectToDB()
	if err := db.Model(&spiel).
		Where("video_id = ?", videoID).
		Column("spiel.*").
		Relation("User").
		Relation("Question").
		Relation("Question.User").
		Select(); err != nil {
		time.Sleep(5 * time.Second)

		return CheckForSpiel(videoID)
	}

	return spiel, nil
}

func UpdateSpielWithVideoURL(spiel models.Spiel) error {
	db = connectToDB()
	_, err := db.Model(&spiel).
		Column("video_url", "created_time").
		WherePK().
		Update()

	if err != nil {
		return err
	}

	return nil
}

func InsertNotificationForSpiel(notification models.Notification) error {
	db = connectToDB()

	if err := db.Insert(&notification); err != nil {
		return err
	}

	return nil
}
