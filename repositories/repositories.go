package repositories

import (
	"discusiin/models"
)

type IDatabase interface {
	SaveNewUser(models.User) error
	Login(username string, password string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

type ITopicDatabase interface {
	GetTopicByName(name string) (models.Topic, error)
	SaveNewTopic(models.Topic) error

	SaveNewModerator(userId int, topicId uint) error
}

type IModeratorDatabase interface {
}
