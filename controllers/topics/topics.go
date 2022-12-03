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
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	//is title empty
	if topic.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "topic name should not be empty")
	}

	//is description empty
	if topic.Description == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "description name should not be empty")
	}

	err := h.ITopicServices.CreateTopic(topic, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Topic created",
	})
}

func (h *TopicHandler) GetAllTopics(c echo.Context) error {
	topics, err := h.ITopicServices.GetTopics()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    topics,
	})
}

func (h *TopicHandler) GetTopic(c echo.Context) error {

	id, errAtoi := strconv.Atoi(c.Param("topic_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	topic, err := h.ITopicServices.GetTopic(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    topic,
	})
}

func (h *TopicHandler) UpdateTopicDescription(c echo.Context) error {

	// validation
	var newTopic models.Topic
	errBind := c.Bind(&newTopic)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, errBind.Error())
	}

	token, errDecodeJWT := helper.DecodeJWT(c)
	if errDecodeJWT != nil {
		return errDecodeJWT
	}

	id, errAtoi := strconv.Atoi(c.Param("topic_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	newTopic.ID = uint(id)

	err := h.ITopicServices.UpdateTopicDescription(newTopic, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Topic updated",
	})
}

func (h *TopicHandler) DeleteTopic(c echo.Context) error {
	id, errAtoi := strconv.Atoi(c.Param("topic_id"))
	if errAtoi != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errAtoi.Error())
	}

	err := h.ITopicServices.RemoveTopic(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Topic deleted",
	})
}
