package models

import (
	"gorm.io/gorm"
)

type FollowedPost struct {
	gorm.Model
	UserID int `json:"userId" form:"userId"`
	PostID int `json:"postId" form:"postId"`

	User User
	Post Post
}

func (FollowedPost) TableName() string {
	return "followedPosts"
}
