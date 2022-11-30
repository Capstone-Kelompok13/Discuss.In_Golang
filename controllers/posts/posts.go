package posts

import (
	"discusiin/helper"
	"discusiin/models"
	"discusiin/services/posts"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	posts.IPostServices
}

func (h *PostHandler) CreateNewPost(c echo.Context) error {
	var p models.Post
	errBind := c.Bind(&p)
	if errBind != nil {
		return errBind
	}

	url_param_value := c.Param("topic_name")
	topicName := helper.URLMinusToSpace(url_param_value)

	//get logged userId
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	err := h.IPostServices.CreatePost(p, topicName, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "post created",
	})
}

func (h *PostHandler) GetAllPost(c echo.Context) error {
	url_param_value := c.Param("topic_name")
	topicName := helper.URLMinusToSpace(url_param_value)

	posts, err := h.IPostServices.GetPosts(topicName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    posts,
	})
}

func (h *PostHandler) GetPost(c echo.Context) error {

	id, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}

	p, err := h.IPostServices.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    p,
	})
}

func (h *PostHandler) EditPost(c echo.Context) error {
	var newPost models.Post
	errBind := c.Bind(&newPost)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errBind.Error(),
		})
	}

	//get user id from logged user
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDecodeJWT.Error(),
		})
	}

	id, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}

	err := h.IPostServices.UpdatePost(newPost, id, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "post updated",
	})
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	var newPost models.Post
	c.Bind(&newPost)

	//get user id from logged user
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDecodeJWT.Error(),
		})
	}

	postID, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}
	err := h.IPostServices.DeletePost(postID, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "post deleted",
	})
}

func (h *PostHandler) GetRecentPost(c echo.Context) error {
	posts, err := h.IPostServices.GetRecentPost()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    posts,
	})
}
