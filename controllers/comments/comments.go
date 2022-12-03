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

// CreateComment godoc
// @Summary Create Comment.
// @Description Create a Comment.
// @Tags Comments
// @Accept json
// @Produce json
// @Security jwt
// @Param topic_id path int true "Topic ID"
// @Param body body map[string]interface{} true "Comment Body"
// @Success 201 {object} map[string]interface{} "comment created"
// @Failure 415 {object} map[string]interface{} "Unsupported Media Type"
// @Failure 400 {object} map[string]interface{} "authorization header not found"
// @Failure 401 {object} map[string]interface{} "token is wrong or expired"
// @Failure 400 {object} map[string]interface{} "error parsing"
// @Failure 404 {object} map[string]interface{} "post not found"
// @Failure 500 {object} map[string]interface{} "error get post by id"
// @Router /api/v1/posts/comments/create/{post_id} [post]
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
		"message": "success",
		"data":    comments,
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
