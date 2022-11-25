package posts

import (
	"discusiin/models"
	"discusiin/repositories"
)

func NewPostServices(db repositories.IDatabase) IPostServices {
	return &postServices{IDatabase: db}
}

type IPostServices interface {
	GetTopic(name string) (int, error)

	CreatePost(post models.Post) error
	SeePosts(id int) ([]models.Post, error)
}

type postServices struct {
	repositories.IDatabase
}

func (p *postServices) GetTopic(name string) (int, error) {
	topic, err := p.IDatabase.GetTopicByName(name)
	if err != nil {
		return 0, err
	}

	return int(topic.ID), nil
}

func (p *postServices) CreatePost(post models.Post) error {
	err := p.IDatabase.SaveNewPost(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) SeePosts(id int) ([]models.Post, error) {
	posts, err := p.IDatabase.GetAllPostByTopic(id)
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}
