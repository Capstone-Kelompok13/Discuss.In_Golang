package dto

import "gorm.io/gorm"

type PublicComment struct {
	gorm.Model
	PostID int         `json:"postId" form:"postId"`
	Body   string      `json:"body" form:"body"`
	User   CommentUser `json:"user" form:"user"`
}

type CommentUser struct {
	UserID   int    `json:"userId" form:"userId"`
	Username string `json:"username" form:"username"`
}
