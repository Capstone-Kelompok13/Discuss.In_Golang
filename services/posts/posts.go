package posts

import (
	"discusiin/models"
	"discusiin/repositories"
	"errors"
	"time"
)

func NewPostServices(db repositories.IDatabase) IPostServices {
	return &postServices{IDatabase: db}
}

type IPostServices interface {
	GetTopic(name string) (int, error)

	CreatePost(post models.Post, name string) error
	SeePosts(name string) ([]models.Post, error)
	SeePost(id int) (models.Post, error)
	UpdatePost(newPost models.Post, id int, userId int) error
	DeletePost(id int, userId int) error
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

func (p *postServices) CreatePost(post models.Post, name string) error {
	//find topic
	topic, err := p.IDatabase.GetTopicByName(name)
	if err != nil {
		return err
	}

	//insert topic id and is active
	post.TopicID = int(topic.ID)
	post.IsActive = true

	//epoch time
	post.CreatedAt = int(time.Now().UnixMilli())

	//save post
	err = p.IDatabase.SaveNewPost(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) SeePosts(name string) ([]models.Post, error) {
	//find topic
	topic, err := p.IDatabase.GetTopicByName(name)
	if err != nil {
		return []models.Post{}, err
	}

	posts, err := p.IDatabase.GetAllPostByTopic(int(topic.ID))
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}

func (p *postServices) SeePost(id int) (models.Post, error) {
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (p *postServices) UpdatePost(newPost models.Post, id int, userId int) error {
	//get previous post
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return err
	}

	if userId != post.UserID {
		return errors.New("user not eligible")
	}

	//update post body
	post.Body += " "
	post.Body += newPost.Body

	err = p.IDatabase.SavePost(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) DeletePost(id int, userId int) error {
	//find post
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return err
	}

	//check user
	if userId != post.UserID {
		return errors.New("user not eligible")
	}

	err = p.IDatabase.DeletePost(id)
	if err != nil {
		return err
	}

	return nil
}
