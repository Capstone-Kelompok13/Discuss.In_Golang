package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Photo    string `json:"photo" form:"photo"`
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin"`
	BanUntil int    `json:"banUntil" form:"banUntil"`
}

func (User) TableName() string {
	return "Users"
}
