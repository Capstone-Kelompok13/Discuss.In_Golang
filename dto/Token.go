package dto

import (
	"github.com/golang-jwt/jwt"
)

type Token struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	jwt.StandardClaims
}
