package models

import (
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	UserID int `json:"userId" form:"userId"`
	PostID int `json:"postId" form:"postId"`

	// User User
	Post Post
}

func (Bookmark) TableName() string {
	return "bookmarks"
}
