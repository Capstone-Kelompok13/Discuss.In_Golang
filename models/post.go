package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string    `json:"title" form:"title"`
	Photo     string    `json:"photo" form:"photo"`
	Body      string    `json:"body" form:"body"`
	UserID    int       `json:"userId" form:"userId"`
	TopicID   int       `json:"topicId" form:"topicId"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	IsActive  bool      `json:"isActive" form:"isActive"`

	Comments []Comment
	User     User
	Topic    Topic
}

func (Post) TableName() string {
	return "posts"
}