package dto

import "gorm.io/gorm"

type PublicReply struct {
	gorm.Model
	CommentID int       `json:"commentId" form:"commentId"`
	Body      string    `json:"body" form:"body"`
	User      ReplyUser `json:"user" form:"user"`
}

type ReplyUser struct {
	UserID   int    `json:"userId" form:"userId"`
	Username string `json:"username" form:"username"`
}
