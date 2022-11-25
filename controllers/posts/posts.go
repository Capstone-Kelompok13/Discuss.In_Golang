package posts

import (
	"discusiin/models"
	"discusiin/services/posts"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	posts.IPostServices
}

func (h *PostHandler) CreateNewPost(c echo.Context) error {
	var p models.Post
	c.Bind(&p)

	name := c.Param("name")

	//get logged userId
	// code here
	p.UserID = 1 //untuk percobaan

	//find topic
	id, err := h.IPostServices.GetTopic(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	p.TopicID = id

	//init field isActive
	p.IsActive = true

	err = h.IPostServices.CreatePost(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "post created",
	})
}

func (h *PostHandler) SeeAllPost(c echo.Context) error {
	var posts []models.Post
	name := c.Param("name")

	//find topic
	id, err := h.IPostServices.GetTopic(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// return c.JSON(http.StatusCreated, map[string]interface{}{
	// 	"message": "success",
	// 	"data":    id,
	// })

	posts, err = h.IPostServices.SeePosts(id)
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
