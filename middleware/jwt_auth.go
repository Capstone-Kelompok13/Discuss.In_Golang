package middleware

import (
	"discusiin/configs"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetToken(id uint, username string) (string, error) {

	log.Println(id, username)
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Cfg.TokenSecret))
}
