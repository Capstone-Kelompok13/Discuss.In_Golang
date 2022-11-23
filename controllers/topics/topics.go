package topics

import (
	"discusiin/models"
	"discusiin/services/topics"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TopicHandler struct {
	topics.ITopicServices
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
