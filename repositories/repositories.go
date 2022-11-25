package repositories

import (
	"discusiin/models"
)

type IDatabase interface {
	SaveNewUser(models.User) error
	Login(username string, password string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)

	GetAllTopics() ([]models.Topic, error)
	GetTopicByName(name string) (models.Topic, error)
	GetTopicByID(id int) (models.Topic, error)
	SaveNewTopic(models.Topic) error
	SaveTopic(models.Topic) error
	RemoveTopic(id int) error
}
