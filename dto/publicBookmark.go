package dto

import (
	"gorm.io/gorm"
)

type PublicBookmark struct {
	gorm.Model
	User BookmarkUser `json:"user" form:"user"`
	Post BookmarkPost `json:"post" form:"post"`
}
type BookmarkPost struct {
	PostID int    `json:"postId" form:"postId"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}
type BookmarkUser struct {
	UserID   int    `json:"userId" form:"userId"`
	Username string `json:"username" form:"username"`
}
