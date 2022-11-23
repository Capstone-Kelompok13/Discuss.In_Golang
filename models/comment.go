package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID     int    `json:"userId" form:"userId"`
	PostID     int    `json:"postId" form:"postId"`
	Body       string `json:"body" form:"body"`
	IsFollowed string `json:"isFollowed" form:"isFollowed"`

	User User
	Post Post
}

func (Comment) TableName() string {
	return "comments"
}
