package topics

import (
	"discusiin/helper"
	"discusiin/models"
	"discusiin/services/topics"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TopicHandler struct {
	topics.ITopicServices
}

func (h *TopicHandler) CreateNewTopic(c echo.Context) error {
	// validation
	var topic models.Topic

	errBind := c.Bind(&topic)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errBind.Error(),
		})
	}

	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDecodeJWT.Error(),
		})
	}

	//is title empty
	if topic.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "topic name should not be empty",
		})
	}

	//is description empty
	if topic.Description == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "topic description should not be empty",
		})
	}

	err := h.ITopicServices.CreateTopic(topic, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "topic created",
	})
}

func (h *TopicHandler) GetAllTopics(c echo.Context) error {
	topics, err := h.ITopicServices.GetTopics()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success",
		"data":    topics,
	})
}

func (h *TopicHandler) GetTopic(c echo.Context) error {

	id, errAtoi := strconv.Atoi(c.Param("topic_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}

	topic, err := h.ITopicServices.GetTopic(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    topic,
	})
}

func (h *TopicHandler) UpdateTopicDescription(c echo.Context) error {

	// validation
	var newTopic models.Topic
	errBind := c.Bind(&newTopic)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errBind.Error(),
		})
	}

	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errDecodeJWT.Error(),
		})
	}

	id, errAtoi := strconv.Atoi(c.Param("topic_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}

	//check if data exist
	topic, err := h.ITopicServices.GetTopic(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	//save topic
	topic.Description = newTopic.Description
	err = h.ITopicServices.SaveTopic(topic, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "topic updated",
	})
}

func (h *TopicHandler) DeleteTopic(c echo.Context) error {
	id, errAtoi := strconv.Atoi(c.Param("topic_id"))
	if errAtoi != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errAtoi.Error(),
		})
	}

	err := h.ITopicServices.RemoveTopic(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "topic deleted",
	})
}
