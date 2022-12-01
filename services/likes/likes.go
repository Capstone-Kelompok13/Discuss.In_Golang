package likes

import (
	"discusiin/dto"
	"discusiin/models"
	"discusiin/repositories"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewLikeServices(db repositories.IDatabase) ILikeServices {
	return &likeServices{IDatabase: db}
}

type ILikeServices interface {
	LikePost(token dto.Token, postId int) error
	DislikePost(token dto.Token, postId int) error
}

type likeServices struct {
	repositories.IDatabase
}

func (l *likeServices) LikePost(token dto.Token, postId int) error {
	var like models.Like
	//cek jika like ada
	oldLike, err := l.IDatabase.GetLikeByUserAndPostId(int(token.ID), postId)
	if err != nil {
		if err.Error() == "record not found" {
			//jika tidak ada
			like.UserID = int(token.ID)
			like.PostID = postId
			like.IsLike = true

			//simpan data like baru
			errSaveLike := l.IDatabase.SaveNewLike(like)
			if errSaveLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errSaveLike.Error())
			}
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		//jika ada
		if oldLike.IsDislike { //jika sudah di like
			//simpan data baru
			oldLike.IsDislike = false //in case like di klik lagi
			oldLike.IsLike = true
			log.Println(oldLike)
			errLike := l.IDatabase.SaveLike(oldLike)
			if errLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errLike.Error())
			}
		}
	}

	return nil
}

func (l *likeServices) DislikePost(token dto.Token, postId int) error {

	var like models.Like

	//cek jika like ada
	oldLike, err := l.IDatabase.GetLikeByUserAndPostId(int(token.ID), postId)
	if err != nil {
		if err.Error() == "record not found" {
			//jika tidak ada
			like.UserID = int(token.ID)
			like.PostID = postId
			like.IsDislike = true

			//simpan data like baru
			errSaveLike := l.IDatabase.SaveNewLike(like)
			if errSaveLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errSaveLike.Error())
			}
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		//jika ada
		if oldLike.IsLike { //jika sudah di dislike
			//simpan data baru
			oldLike.IsDislike = true //in case like di klik lagi supaya netral
			oldLike.IsLike = false
			errLike := l.IDatabase.SaveLike(oldLike)
			if errLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errLike.Error())
			}
		}
	}

	return nil
}
