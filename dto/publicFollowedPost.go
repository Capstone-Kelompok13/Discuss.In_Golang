package dto

import (
	"gorm.io/gorm"
)

type PublicFollowedPost struct {
	gorm.Model
	User FollowedPostUser `json:"user" form:"user"`
	Post FollowedPost     `json:"post" form:"post"`
}
type FollowedPost struct {
	PostID int    `json:"postId" form:"postId"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}
type FollowedPostUser struct {
	UserID   int    `json:"userId" form:"userId"`
	Username string `json:"username" form:"username"`
}
