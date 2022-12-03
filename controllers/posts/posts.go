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
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	url_param_value := c.Param("topic_name")
	topicName := helper.URLDecodeReformat(url_param_value)

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
		"message": "Post created",
	})
}

func (h *PostHandler) GetAllPost(c echo.Context) error {
	url_param_value := c.Param("topic_name")
	topicName := helper.URLDecodeReformat(url_param_value)
	if url_param_value == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "topic name should not be empty")
	}
	if topicName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "topic name should not be empty")
	}
	//check if page exist
	page, errAtoi := strconv.Atoi(c.QueryParam("page"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	posts, numberOfPage, err := h.IPostServices.GetPosts(topicName, page)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "Success",
		"data":           posts,
		"number_of_page": numberOfPage,
		"page":           page,
	})
}

func (h *PostHandler) GetPost(c echo.Context) error {

	id, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	p, err := h.IPostServices.GetPost(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    p,
	})
}

func (h *PostHandler) EditPost(c echo.Context) error {
	var newPost models.Post
	errBind := c.Bind(&newPost)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	//get user id from logged user
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	id, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	err := h.IPostServices.UpdatePost(newPost, id, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post updated",
	})
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	var newPost models.Post
	c.Bind(&newPost)

	//get user id from logged user
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	postID, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}
	err := h.IPostServices.DeletePost(postID, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post deleted",
	})
}

func (h *PostHandler) GetRecentPost(c echo.Context) error {
	//check if page exist
	page, _ := strconv.Atoi(c.QueryParam("page"))

	posts, numberOfPage, err := h.IPostServices.GetRecentPost(page)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":        "Success",
		"data":           posts,
		"number_of_page": numberOfPage,
		"page":           page,
	})
}
