package posts

import (
	"discusiin/models"
	"discusiin/repositories"
	"errors"
)

func NewCommentServices(db repositories.IDatabase) ICommentServices {
	return &commentServices{IDatabase: db}
}

type ICommentServices interface {
	CreateComment(comment models.Comment, id int) error
	SeeAllComments(id int) ([]models.Comment, error)
	UpdateComment(newComment models.Comment, id int, co int, userId int) error
	DeleteComment(userId int, co int) error
}

type commentServices struct {
	repositories.IDatabase
}

func (c *commentServices) CreateComment(comment models.Comment, id int) error {
	//get post
	post, err := c.IDatabase.GetPostById(id)
	if err != nil {
		return err
	}

	//fill comment empty comment field
	comment.PostID = int(post.ID)
	comment.IsFollowed = true

	err = c.IDatabase.SaveNewComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (c *commentServices) SeeAllComments(id int) ([]models.Comment, error) {
	comments, err := c.IDatabase.GetAllCommentByPost(id)
	if err != nil {
		return []models.Comment{}, err
	}

	return comments, nil
}

func (c *commentServices) UpdateComment(newComment models.Comment, id int, co int, userId int) error {
	//get comment
	comment, err := c.IDatabase.GetCommentById(co)
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
	err = c.IDatabase.SaveComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (c *commentServices) DeleteComment(userId int, co int) error {
	//get comment
	comment, err := c.IDatabase.GetCommentById(co)
	if err != nil {
		return err
	}

	//check user
	if comment.UserID != userId {
		return errors.New("user not eligible")
	}

	err = c.IDatabase.DeleteComment(co)
	if err != nil {
		return err
	}

	return nil
}
