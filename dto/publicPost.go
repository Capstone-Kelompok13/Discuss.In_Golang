package dto

import "gorm.io/gorm"

type PublicPost struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Photo     string `json:"photo" form:"photo"`
	Body      string `json:"body" form:"body"`
	UserID    int    `json:"userId" form:"userId"`
	TopicID   int    `json:"topicId" form:"topicId"`
	CreatedAt int    `json:"createdAt" form:"createdAt"`
	IsActive  bool   `json:"isActive" form:"isActive"`
}
