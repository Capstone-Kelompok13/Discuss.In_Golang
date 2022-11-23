package models

import (
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	UserID    int    `json:"userId" form:"userId"`
	CommentID int    `json:"commentId" form:"commentId"`
	Body      string `json:"body" form:"body"`

	User    User
	Comment Comment
}

func (Reply) TableName() string {
	return "Replys"
}
