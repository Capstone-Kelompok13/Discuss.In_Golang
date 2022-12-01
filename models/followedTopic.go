package models

import (
	"gorm.io/gorm"
)

type FollowedTopic struct {
	gorm.Model
	UserID  int `json:"userId" form:"userId"`
	TopicID int `json:"topicId" form:"topicId"`

	User  User
	Topic Topic
}

func (FollowedTopic) TableName() string {
	return "followedTopics"
}
