package replies

import (
	"discusiin/helper"
	"discusiin/models"
	"discusiin/services/replies"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReplyHandler struct {
	replies.IReplyServices
}

func (h *ReplyHandler) CreateReply(c echo.Context) error {
	var reply models.Reply
	// c.Bind(&reply)
	errBind := c.Bind(&reply)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	//get logged userId
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	//get comment id
	commentId, errAtoi := strconv.Atoi(c.Param("comment_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	err := h.IReplyServices.CreateReply(reply, commentId, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Reply created",
	})
}

func (h *ReplyHandler) GetAllReply(c echo.Context) error {
	//get comment id
	commentId, errAtoi := strconv.Atoi(c.Param("comment_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	//get all reply from comment
	replies, err := h.IReplyServices.GetAllReply(commentId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    replies,
	})
}

func (h *ReplyHandler) UpdateReply(c echo.Context) error {
	var newReply models.Reply
	errBind := c.Bind(&newReply)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	//get logged userId
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	//get reply id
	replyId, errAtoi := strconv.Atoi(c.Param("reply_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	//update reply
	err := h.IReplyServices.UpdateReply(newReply, replyId, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Reply updated",
	})
}

func (h *ReplyHandler) DeleteReply(c echo.Context) error {
	//get logged userId
	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	//get reply id
	replyId, errAtoi := strconv.Atoi(c.Param("reply_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	//delete reply
	err := h.IReplyServices.DeleteReply(replyId, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Reply deleted",
	})
}
