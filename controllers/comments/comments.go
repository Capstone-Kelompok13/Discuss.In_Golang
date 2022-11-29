package comments

import (
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
	c.Bind(&comment)

	//get logged userId
	// code here
	comment.UserID = 1 //untuk percobaan

	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ICommentServices.CreateComment(comment, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Comment Created",
	})

}

func (h *CommentHandler) SeeAllComment(c echo.Context) error {
	var comments []models.Comment

	//get topic id
	id, _ := strconv.Atoi(c.Param("id"))

	//get all coment from post
	comments, err := h.ICommentServices.SeeAllComments(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Succes",
		"data":    comments,
	})
}

func (h *CommentHandler) UpdateComment(c echo.Context) error {
	var comment models.Comment
	c.Bind(&comment)

	//get logged userId
	// code here
	userId := 1 //untuk percobaan

	//check if user who eligible
	co, _ := strconv.Atoi(c.Param("co")) //untuk param comment
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ICommentServices.UpdateComment(comment, id, co, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Comment Updated",
	})

}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	//get logged userId
	// code here
	userId := 1 //untuk percobaan

	//check if user who eligible
	co, _ := strconv.Atoi(c.Param("co")) //untuk param comment
	err := h.ICommentServices.DeleteComment(userId, co)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Comment Deleted",
	})
}
