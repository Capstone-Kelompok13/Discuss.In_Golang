package users

import (
	"discusiin/dto"
	"discusiin/middleware"
	"discusiin/models"
	"discusiin/repositories"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewUserServices(db repositories.IDatabase) IUserServices {
	return &userServices{IDatabase: db}
}

type IUserServices interface {
	Register(user models.User) error
	Login(user models.User) (dto.Login, error)
}

type userServices struct {
	repositories.IDatabase
}

func (s *userServices) Register(user models.User) error {

	_, err1 := s.IDatabase.GetUserByUsername(user.Username)
	if err1 != nil {
		if err1.Error() == "record not found" {
			err2 := s.IDatabase.SaveNewUser(user)
			if err2 != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err2.Error())
			}
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err1.Error())
		}
	} else {
		// if getUserByUsername got no error
		return errors.New("username exist, try another username")
	}
	return nil
}
func (s *userServices) Login(user models.User) (dto.Login, error) {

	user, err := s.IDatabase.Login(user.Email, user.Password)
	if err != nil {
		if err.Error() == "record not found" {
			return dto.Login{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return dto.Login{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := middleware.GetToken(user.ID, user.Username)
	if err != nil {
		return dto.Login{}, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	result := dto.Login{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Photo:    user.Photo,
		BanUntil: user.BanUntil,
		IsAdmin:  user.IsAdmin,
		Token:    token,
	}

	return result, nil
}
