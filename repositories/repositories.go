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

	SaveNewPost(post models.Post) error
	GetAllPostByTopic(id int) ([]models.Post, error)
	GetPostById(id int) (models.Post, error)
	SavePost(post models.Post) error
	DeletePost(id int) error
	GetPostByIdWithAll(id int) (models.Post, error)

	SaveNewComment(comment models.Comment) error
	GetCommentById(co int) (models.Comment, error)
	SaveComment(comment models.Comment) error
	DeleteComment(co int) error
}
