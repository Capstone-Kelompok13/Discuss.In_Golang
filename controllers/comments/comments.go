package comments

import (
	"discusiin/helper"
	"discusiin/models"
	comments "discusiin/services/comments"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	comments.ICommentServices
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	var comment models.Comment
	errBind := c.Bind(&comment)
	if errBind != nil {
		echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	//get logged userId
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	postID, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	err := h.ICommentServices.CreateComment(comment, postID, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "comment created",
	})

}

func (h *CommentHandler) GetAllComment(c echo.Context) error {

	//get post id
	postID, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}
	//get all coment from post
	comments, err := h.ICommentServices.GetAllComments(postID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success",
		"data_comment": comments,
	})
}

func (h *CommentHandler) UpdateComment(c echo.Context) error {
	var comment models.Comment
	errBind := c.Bind(&comment)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	//get logged userId
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	//check if user who eligible untuk param comment
	commentID, errAtoi := strconv.Atoi(c.Param("comment_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	comment.ID = uint(commentID)

	err := h.ICommentServices.UpdateComment(comment, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "comment updated",
	})

}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	//get logged userId
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	//check if user who eligible
	commentID, errAtoi := strconv.Atoi(c.Param("comment_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}
	err := h.ICommentServices.DeleteComment(commentID, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "comment deleted",
	})
}
