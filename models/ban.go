package models

import (
	"time"

	"gorm.io/gorm"
)

type Ban struct {
	gorm.Model
	UserID  int       `json:"userId" form:"userId"`
	TopicID int       `json:"topicId" form:"topicId"`
	Ban     time.Time `json:"ban" form:"ban"`

	User  User
	Topic Topic
}

func (Ban) TableName() string {
	return "bans"
}
