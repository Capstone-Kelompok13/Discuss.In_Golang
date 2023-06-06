package dto

type Login struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Photo    string `json:"photo" form:"photo"`
	Token    string `json:"token" form:"token"`
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin"`
	BanUntil int    `json:"banUntil" form:"banUntil"`
}
