package likes

import (
	"discusiin/helper"
	"discusiin/models"
	"discusiin/services/likes"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	likes.ILikeServices
}

func (h *LikeHandler) LikePost(c echo.Context) error {
	var like models.Like

	errBind := c.Bind(&like)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errBind.Error(),
		})
	}

	//get logged userid
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDecodeJWT.Error(),
		})
	}

	//get post id
	postId, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}

	err := h.ILikeServices.LikePost(token, postId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Like Success",
	})
}

func (h *LikeHandler) DislikePost(c echo.Context) error {
	var like models.Like

	errBind := c.Bind(&like)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errBind.Error(),
		})
	}

	//get logged userid
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDecodeJWT.Error(),
		})
	}

	//get post id
	postId, errAtoi := strconv.Atoi(c.Param("post_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}

	err := h.ILikeServices.DislikePost(token, postId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Like Success",
	})
}
