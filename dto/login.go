package dto

type Login struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}
