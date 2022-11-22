package middleware

import (
	"discusiin/configs"
	"discusiin/dto"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetToken(id uint, username string) (string, error) {

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Cfg.TokenSecret))
}

func DecodeJWT(ctx echo.Context) (dto.Token, error) {
	var t dto.Token

	auth := ctx.Request().Header.Get("Authorization")
	if auth == "" {
		return dto.Token{}, errors.New("authorization header not found")
	}

	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]

	token, err := jwt.ParseWithClaims(auth, &dto.Token{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.Cfg.TokenSecret), nil
	})
	if err != nil {
		return dto.Token{}, errors.New("token is wrong or expired")
	}

	if claims, ok := token.Claims.(*dto.Token); ok && token.Valid {
		t.ID = claims.ID
		t.Username = claims.Username
	}

	return t, nil
}
