package helper

import (
	"discusiin/configs"
	"discusiin/dto"
	"errors"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
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

func URLMinusToSpace(url_param_value string) string {
	return strings.ReplaceAll(url_param_value, "-", " ")
}
