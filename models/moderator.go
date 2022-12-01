package models

import (
	"gorm.io/gorm"
)

type Moderator struct {
	gorm.Model
	UserID  int `json:"userId" form:"userId"`
	TopicID int `json:"topicId" form:"topicId"`

	User  User
	Topic Topic
}

func (Moderator) TableName() string {
	return "Moderators"
}
