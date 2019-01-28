package database

import "spiel/notification-center/models"

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
