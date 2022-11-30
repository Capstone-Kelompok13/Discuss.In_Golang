package posts

import (
	"discusiin/dto"
	"discusiin/models"
	"discusiin/repositories"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NewPostServices(db repositories.IDatabase) IPostServices {
	return &postServices{IDatabase: db}
}

type IPostServices interface {
	GetTopic(name string) (int, error)

	CreatePost(post models.Post, name string, token dto.Token) error
	GetPosts(name string) ([]dto.PublicPost, error)
	GetPost(id int) (dto.PublicPost, error)
	UpdatePost(newPost models.Post, id int, token dto.Token) error
	DeletePost(id int, token dto.Token) error
	GetRecentPost() ([]dto.PublicPost, error)
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

func (p *postServices) CreatePost(post models.Post, name string, token dto.Token) error {
	//find topic
	topic, err := p.IDatabase.GetTopicByName(name)
	if err != nil {
		return err
	}

	//owner
	post.UserID = int(token.ID)
	//insert topic id and is active
	post.TopicID = int(topic.ID)
	//epoch time
	post.CreatedAt = int(time.Now().UnixMilli())
	// isActiveDefault
	post.IsActive = true

	//save post
	err = p.IDatabase.SaveNewPost(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) GetPosts(name string) ([]dto.PublicPost, error) {
	//find topic
	topic, err := p.IDatabase.GetTopicByName(name)
	if err != nil {
		return nil, err
	}

	posts, err := p.IDatabase.GetAllPostByTopic(int(topic.ID))
	var result []dto.PublicPost
	for _, v := range posts {
		result = append(result, dto.PublicPost{
			Model:     v.Model,
			Title:     v.Title,
			Photo:     v.Photo,
			Body:      v.Body,
			UserID:    v.UserID,
			Username:  v.User.Username,
			TopicID:   v.TopicID,
			Topicname: v.Topic.Name,
			CreatedAt: v.CreatedAt,
			IsActive:  v.IsActive,
		})
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *postServices) GetPost(id int) (dto.PublicPost, error) {
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return dto.PublicPost{}, err
	}
	result := dto.PublicPost{
		Model:     post.Model,
		Title:     post.Title,
		Photo:     post.Photo,
		Body:      post.Body,
		UserID:    post.UserID,
		Username:  post.User.Username,
		TopicID:   post.TopicID,
		Topicname: post.Topic.Name,
		CreatedAt: post.CreatedAt,
		IsActive:  post.IsActive,
	}

	return result, nil
}

func (p *postServices) UpdatePost(newPost models.Post, postID int, token dto.Token) error {
	//get previous post
	post, err := p.IDatabase.GetPostById(postID)
	if err != nil {
		return err
	}

	if int(token.ID) != post.UserID {
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

func (p *postServices) DeletePost(id int, token dto.Token) error {
	//find post
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return err
	}

	//check user
	user, err := p.IDatabase.GetUserByUsername(token.Username)
	if err != nil {
		return err
	}

	if !user.IsAdmin {
		if int(token.ID) != post.UserID {
			return errors.New("user not eligible")
		}
	}

	err = p.IDatabase.DeletePost(id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) GetRecentPost() ([]dto.PublicPost, error) {
	posts, err := p.IDatabase.GetRecentPost()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var result []dto.PublicPost
	for _, post := range posts {
		result = append(result, dto.PublicPost{
			Model:     post.Model,
			Title:     post.Title,
			Photo:     post.Photo,
			Body:      post.Body,
			UserID:    post.UserID,
			Username:  post.User.Username,
			TopicID:   post.TopicID,
			Topicname: post.Topic.Name,
			CreatedAt: post.CreatedAt,
			IsActive:  post.IsActive,
		})
	}

	return result, nil
}
