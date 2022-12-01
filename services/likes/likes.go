package likes

import (
	"discusiin/dto"
	"discusiin/models"
	"discusiin/repositories"
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
		if oldLike.IsLike && !oldLike.IsDislike { //jika sudah di like
			//simpan data baru
			oldLike.IsLike = false //in case like di klik lagi

			errLike := l.IDatabase.SaveLike(oldLike)
			if errLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errLike.Error())
			}
		} else if oldLike.IsDislike { //jika di dislike
			//simpan data baru
			oldLike.IsDislike = false
			oldLike.IsLike = true

			errLike := l.IDatabase.SaveLike(oldLike)
			if errLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errLike.Error())
			}
		} else { //jika like netral
			oldLike.IsLike = true

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
		//jika ada
		if !oldLike.IsLike && oldLike.IsDislike { //jika sudah di dislike
			//simpan data baru
			oldLike.IsDislike = false //in case like di klik lagi supaya netral

			errLike := l.IDatabase.SaveLike(oldLike)
			if errLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errLike.Error())
			}
		} else if oldLike.IsLike { //jika di like
			//simpan data baru
			oldLike.IsDislike = true
			oldLike.IsLike = false

			errLike := l.IDatabase.SaveLike(oldLike)
			if errLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errLike.Error())
			}
		} else { //jika like netral
			oldLike.IsDislike = true

			errLike := l.IDatabase.SaveLike(oldLike)
			if errLike != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errLike.Error())
			}
		}
	}

	return nil
}
