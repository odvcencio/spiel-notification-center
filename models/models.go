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
	Categories        []map[string]interface{} `json:"user_categories,omitempty" sql:"-"`
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
	tableName  struct{}  `sql:"user_categories,alias:user_category"`
	UserID     string    `json:"-"`
	User       *User     `json:"-" sql:",fk"`
	CategoryID int       `json:"-"`
	Category   *Category `json:"category,omitempty" sql:",fk"`
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
	User             *User     `json:"asker,omitempty" sql:",fk"`
	Question         string    `json:"question"`
	Category         *Category `json:"category,omitempty" sql:",fk"`
	CategoryID       int       `json:"-"`
	SpielsCount      int       `json:"spiels_count"`
	ID               int       `json:"id"`

	// Relations
	Spiels   []Spiel  `json:"spiels,omitempty"`
	Receiver *AskUser `json:"receiver,omitempty"`
}

type AskUser struct {
	tableName  struct{}  `sql:"ask_user,alias:ask_user"`
	QuestionID int       `json:"-"`
	Question   *Question `json:"question,omitempty" sql:",fk"`
	UserID     string    `json:"-"`
	User       *User     `json:"user,omitempty" sql:",fk"`
}

type Spiel struct {
	User             *User             `json:"spieler,omitempty" sql:",fk"`
	UserID           string            `json:"-"`
	VideoURL         string            `json:"video_url,omitempty"`
	VideoID          string            `json:"video_id,omitempty"`
	Duration         float32           `json:"duration,omitempty"`
	ThumbnailURL     string            `json:"thumbnail_url,omitempty"`
	Question         *Question         `json:"question,omitempty" sql:",fk"`
	QuestionID       int               `json:"-"`
	Category         *Category         `json:"category,omitempty" sql:",fk"`
	CategoryID       int               `json:"-"`
	Assessable       bool              `json:"assessable,omitempty" sql:"-"`
	Assessments      []SpielAssessment `json:"-"`
	LocalCreatedTime string            `json:"local_time,omitempty" sql:"-"`
	CreatedTime      *time.Time        `json:"created_time,omitempty"`
	ID               int               `json:"id"`
}

type SpielAssessment struct {
	ID          int       `json:"id"`
	SpielID     int       `json:"-"`
	Spiel       *Spiel    `json:"-" sql:",fk"`
	UserID      string    `json:"-"`
	User        *User     `json:"-" sql:",fk"`
	CreatedTime time.Time `json:"created_time"`

	// Relations
	Choices []SpielAssessmentChoice `json:"choices,omitempty" pg:"many2many:spiel_assessment_to_choices"`
}

type SpielAssessmentChoice struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Selected bool   `json:"selected" sql:"-"`
}

type SpielAssessmentToChoice struct {
	tableName    struct{}               `sql:"spiel_assessment_to_choices"`
	Assessment   *SpielAssessment       `json:"-" pg:"fk:spiel_assessment_id"`
	AssessmentID int                    `sql:"spiel_assessment_id"`
	Choice       *SpielAssessmentChoice `json:"-" pg:"fk:spiel_assessment_choice_id"`
	ChoiceID     int                    `sql:"spiel_assessment_choice_id"`
}

type Notification struct {
	ID                int              `json:"-"`
	SpielAssessmentID int              `json:"-"`
	SpielAssessment   *SpielAssessment `json:"assessment,omitempty" sql:",fk"`
	UserID            string           `json:"-"`
	Message           string           `json:"message"`
	SpielID           int              `json:"-"`
	Spiel             *Spiel           `json:"spiel,omitempty" sql:",fk"`
	Type              string           `json:"type"`
	CreatedTime       time.Time        `json:"created_time"`
}

func init() {
	orm.RegisterTable((*SpielAssessmentToChoice)(nil))
}
