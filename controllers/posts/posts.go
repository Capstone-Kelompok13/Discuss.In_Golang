package posts

import (
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
	c.Bind(&p)

	name := c.Param("name")

	//get logged userId
	// code here
	p.UserID = 1 //untuk percobaan

	err := h.IPostServices.CreatePost(p, name)
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

func (h *PostHandler) SeePost(c echo.Context) error {
	var p models.Post

	id, _ := strconv.Atoi(c.Param("id"))

	p, err := h.IPostServices.SeePost(id)
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
	c.Bind(&newPost)

	//check if user is correct
	// code

	id, _ := strconv.Atoi(c.Param("id"))
	err := h.IPostServices.UpdatePost(newPost, id)
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

	//check if user is correct
	// code

	id, _ := strconv.Atoi(c.Param("id"))
	err := h.IPostServices.DeletePost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "post deleted",
	})
}

func (h *PostHandler) CreateComment(c echo.Context) error {
	var comment models.Comment
	c.Bind(&comment)

	//get logged userId
	// code here
	comment.UserID = 1 //untuk percobaan

	id, _ := strconv.Atoi(c.Param("id"))
	err := h.IPostServices.CreateComment(comment, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Comment Created",
	})

	// //find post if post exist
	// id, _ := strconv.Atoi(c.Param("id"))
	// post, err := h.IPostServices.SeePost(id)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": err.Error(),
	// 	})
	// }

	// //get logged userId
	// // code here
	// comment.UserID = 1 //untuk percobaan

}

func (h *PostHandler) UpdateComment(c echo.Context) error {
	var comment models.Comment
	c.Bind(&comment)

	//get logged userId
	// code here
	userId := 1 //untuk percobaan

	//check if user who eligible
	co, _ := strconv.Atoi(c.Param("c")) //untuk param comment
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.IPostServices.UpdateComment(comment, id, co, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Comment Updated",
	})

	// //find post if post exist
	// id, _ := strconv.Atoi(c.Param("id"))
	// post, err := h.IPostServices.SeePost(id)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": err.Error(),
	// 	})
	// }

	// //get logged userId
	// // code here
	// comment.UserID = 1 //untuk percobaan

}

func (h *PostHandler) DeleteComment(c echo.Context) error {
	//get logged userId
	// code here
	userId := 1 //untuk percobaan

	//check if user who eligible
	co, _ := strconv.Atoi(c.Param("c")) //untuk param comment
	err := h.IPostServices.DeleteComment(co, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Comment Deleted",
	})
}
