package topics

import (
	"discusiin/models"
	"discusiin/services/topics"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TopicHandler struct {
	topics.ITopicServices
}

func (h *TopicHandler) SeeAllTopics(c echo.Context) error {
	var topics []models.Topic

	topics, err := h.ITopicServices.SeeTopics()
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

func (h *TopicHandler) CreateNewTopic(c echo.Context) error {
	// validation
	var t models.Topic

	err := c.Bind(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	//is title empty
	if t.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "topic name or title should not be empty",
		})
	}

	//is description empty
	if t.Description == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "topic description should not be empty",
		})
	}

	err = h.ITopicServices.CreateTopic(t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "topic created",
	})
}

func (h *TopicHandler) SeeTopic(c echo.Context) error {
	var topic models.Topic

	id, _ := strconv.Atoi(c.Param("id"))

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

func (h *TopicHandler) UpdateDescriptionTopic(c echo.Context) error {
	// validation
	newTopic := models.Topic{}
	c.Bind(&newTopic)

	id, _ := strconv.Atoi(c.Param("id"))

	//check if data exist
	topic, err := h.ITopicServices.GetTopic(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	//save topic
	topic.Description = newTopic.Description
	err = h.ITopicServices.SaveTopic(topic)
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
	id, _ := strconv.Atoi(c.Param("id"))

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
