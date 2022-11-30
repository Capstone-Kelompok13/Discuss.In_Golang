package dto

import "gorm.io/gorm"

type PublicComment struct {
	gorm.Model
	UserID   int    `json:"userId" form:"userId"`
	PostID   int    `json:"postId" form:"postId"`
	Body     string `json:"body" form:"body"`
	Username string `json:"username" form:"username"`
}
