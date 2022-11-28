package posts

import (
	"discusiin/models"
	"discusiin/repositories"
	"errors"
)

func NewPostServices(db repositories.IDatabase) IPostServices {
	return &postServices{IDatabase: db}
}

type IPostServices interface {
	GetTopic(name string) (int, error)

	CreatePost(post models.Post) error
	SeePosts(id int) ([]models.Post, error)
	SeePost(id int) (models.Post, error)
	UpdatePost(newPost models.Post, id int) error
	DeletePost(id int) error

	CreateComment(comment models.Comment, id int) error
	UpdateComment(newComment models.Comment, id int, co int, userId int) error
	DeleteComment(userId int, co int) error
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

func (p *postServices) SeePost(id int) (models.Post, error) {
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (p *postServices) UpdatePost(newPost models.Post, id int) error {
	//get previous post
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return err
	}
	post.Body += " "
	post.Body += newPost.Body

	err = p.IDatabase.SavePost(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) DeletePost(id int) error {
	err := p.IDatabase.DeletePost(id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) CreateComment(comment models.Comment, id int) error {
	//get post
	post, err := p.IDatabase.GetPostById(id)
	if err != nil {
		return err
	}

	//fill comment empty comment field
	comment.PostID = int(post.ID)
	comment.IsFollowed = true

	err = p.IDatabase.SaveNewComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) UpdateComment(newComment models.Comment, id int, co int, userId int) error {
	// //get post
	// post, err := p.IDatabase.GetPostById(id)
	// if err != nil {
	// 	return err
	// }

	//get comment
	comment, err := p.IDatabase.GetCommentById(co)
	if err != nil {
		return err
	}

	//check user
	if comment.UserID != userId {
		return errors.New("user not eligible")
	}

	//update comment field
	comment.Body += " "
	comment.Body += newComment.Body

	//save comment
	err = p.IDatabase.SaveComment(newComment)
	if err != nil {
		return err
	}

	return nil
}

func (p *postServices) DeleteComment(userId int, co int) error {
	//get comment
	comment, err := p.IDatabase.GetCommentById(co)
	if err != nil {
		return err
	}

	//check user
	if comment.UserID != userId {
		return errors.New("user not eligible")
	}

	err = p.IDatabase.DeleteComment(co)
	if err != nil {
		return err
	}

	return nil
}
