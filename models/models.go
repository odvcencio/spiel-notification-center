package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type User struct {
	Company           string                   `json:"company"`
	Email             string                   `json:"email"`
	Username          string                   `json:"username"`
	FirstName         string                   `json:"first_name"`
	LastName          string                   `json:"last_name"`
	Title             string                   `json:"title"`
	ProfilePhotoURL   string                   `json:"profile_photo_url"`
	CreatedTime       time.Time                `json:"-"`
	FriendlyJoinDate  string                   `json:"friendly_join_date" sql:"-"`
	ViewIndex         int                      `json:"view_index"`
	AboutMe           string                   `json:"about_me"`
	PhoneNumber       string                   `json:"phone_number"`
	UserCategories    []UserCategory           `json:"-" sql:",fk"`
	Categories        []map[string]interface{} `json:"user_categories" sql:"-"`
	ID                string                   `json:"id"`
	QuestionsToAnswer []Question               `json:"-" sql:"many2many:ask_user,joinFK:user_id"`
}

func (target *User) Merge(source User) User {
	if target.Username != "" {
		source.Username = target.Username
	}
	if target.FirstName != "" {
		source.FirstName = target.FirstName
	}
	if target.LastName != "" {
		source.LastName = target.LastName
	}
	if target.Title != "" {
		source.Title = target.Title
	}
	if target.Company != "" {
		source.Company = target.Company
	}
	if target.ViewIndex != -1 {
		source.ViewIndex = target.ViewIndex
	}
	if target.AboutMe != "" {
		source.AboutMe = target.AboutMe
	}
	if target.Email != "" {
		source.Email = target.Email
	}
	if target.ID != "" && target.ID != source.Email {
		source.ID = target.ID
	}

	return source
}

type UserCategory struct {
	tableName  struct{} `sql:"user_categories,alias:user_category"`
	UserID     string   `json:"-"`
	User       User     `json:"-" sql:",fk"`
	CategoryID int      `json:"-"`
	Category   Category `json:"category" sql:",fk"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	Answered         bool      `json:"answered"`
	CreatedTime      time.Time `json:"created_time"`
	LocalCreatedTime string    `json:"local_time" sql:"-"`
	UserID           string    `json:"-"`
	User             User      `json:"asker" sql:",fk"`
	Question         string    `json:"question"`
	Category         Category  `json:"category" sql:",fk"`
	CategoryID       int       `json:"category_id"`
	ID               int       `json:"id"`
}

type AskUser struct {
	tableName  struct{} `sql:"ask_user,alias:ask_user"`
	QuestionID int      `json:"-"`
	Question   Question `json:"question" sql:",fk"`
	UserID     string   `json:"-"`
	User       User     `json:"user" sql:",fk"`
}

type Spiel struct {
	User             User              `json:"spieler" sql:",fk"`
	UserID           string            `json:"-"`
	VideoURL         string            `json:"video_url"`
	VideoID          string            `json:"video_id"`
	Question         Question          `json:"question" sql:",fk"`
	QuestionID       int               `json:"-"`
	Category         Category          `json:"category" sql:",fk"`
	CategoryID       int               `json:"-"`
	Assessable       bool              `json:"assessable" sql:"-"`
	Assessments      []SpielAssessment `json:"-"`
	LocalCreatedTime string            `json:"local_time" sql:"-"`
	CreatedTime      time.Time         `json:"created_time"`
	ID               int               `json:"id"`
}

type SpielAssessment struct {
	ID          int       `json:"id"`
	SpielID     int       `json:"-"`
	Spiel       Spiel     `json:"-" sql:",fk"`
	UserID      string    `json:"-"`
	User        User      `json:"-" sql:",fk"`
	CreatedTime time.Time `json:"created_time"`

	// Relations
	Choices []SpielAssessmentChoice `json:"choices" pg:"many2many:spiel_assessment_to_choices"`
}

type SpielAssessmentChoice struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Selected bool   `json:"selected" sql:"-"`
}

type SpielAssessmentToChoice struct {
	tableName    struct{}              `sql:"spiel_assessment_to_choices"`
	Assessment   SpielAssessment       `json:"assessment" pg:"fk:spiel_assessment_id"`
	AssessmentID int                   `sql:"spiel_assessment_id"`
	Choice       SpielAssessmentChoice `json:"choice" pg:"fk:spiel_assessment_choice_id"`
	ChoiceID     int                   `sql:"spiel_assessment_choice_id"`
}

type Notification struct {
	ID                int             `json:"-"`
	UserID            string          `json:"-"`
	Message           string          `json:"message"`
	SpielID           int             `json:"-"`
	Spiel             Spiel           `json:"spiel" sql:",fk"`
	SpielAssessmentID int             `json:"-"`
	SpielAssessment   SpielAssessment `json:"assessment" sql:",fk"`
	Type              string          `json:"type"`
	CreatedTime       time.Time       `json:"created_time"`
}

func init() {
	orm.RegisterTable((*SpielAssessmentToChoice)(nil))
}
