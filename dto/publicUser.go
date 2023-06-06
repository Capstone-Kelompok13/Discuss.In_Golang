package dto

type PublicUser struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Photo    string `json:"photo" form:"photo"`
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin"`
	BanUntil int    `json:"banUntil" form:"banUntil"`
}
