package users

import (
	"discusiin/dto"
	"discusiin/middleware"
	"discusiin/models"
	"discusiin/repositories"
	"errors"
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
				return err2
			}
		} else {
			return err1
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
		return dto.Login{}, err
	}

	token, err := middleware.GetToken(user.ID, user.Username)
	if err != nil {
		return dto.Login{}, err
	}

	var result dto.Login
	result.ID = user.ID
	result.Username = user.Username
	result.Token = token

	return result, nil
}
