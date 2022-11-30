package topics

import (
	"discusiin/dto"
	"discusiin/models"
	"discusiin/repositories"
	"errors"
)

func NewTopicServices(db repositories.IDatabase) ITopicServices {
	return &topicServices{IDatabase: db}
}

type ITopicServices interface {
	GetTopics() ([]models.Topic, error)
	CreateTopic(topic models.Topic, token dto.Token) error
	GetTopic(id int) (models.Topic, error)
	SaveTopic(topic models.Topic, token dto.Token) error
	RemoveTopic(id int) error
}

type topicServices struct {
	repositories.IDatabase
}

func (t *topicServices) GetTopics() ([]models.Topic, error) {
	//get all topics
	topics, err := t.IDatabase.GetAllTopics()
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (t *topicServices) CreateTopic(topic models.Topic, token dto.Token) error {
	// isAdmin?
	user, errGetUser := t.IDatabase.GetUserByUsername(token.Username)
	if errGetUser != nil {
		return errGetUser
	}
	if !user.IsAdmin {
		return errors.New("admin access only")
	}

	// isExist?
	_, err1 := t.IDatabase.GetTopicByName(topic.Name)
	if err1 != nil {
		if err1.Error() == "record not found" {
			err2 := t.IDatabase.SaveNewTopic(topic)
			if err2 != nil {
				return err2
			}
		} else {
			return err1
		}
	} else {
		//if topic exist
		return errors.New("topic already exist")
	}

	return nil
}

func (t *topicServices) GetTopic(id int) (models.Topic, error) {
	topic, err := t.IDatabase.GetTopicByID(id)
	if err != nil {
		return models.Topic{}, err
	}
	return topic, nil
}

func (t *topicServices) SaveTopic(topic models.Topic, token dto.Token) error {
	// isAdmin?
	user, errGetUser := t.IDatabase.GetUserByUsername(token.Username)
	if errGetUser != nil {
		return errGetUser
	}
	if !user.IsAdmin {
		return errors.New("admin access only")
	}

	err := t.IDatabase.SaveTopic(topic)
	if err != nil {
		return err
	}

	return nil
}

func (t *topicServices) RemoveTopic(id int) error {
	err := t.IDatabase.RemoveTopic(id)

	if err != nil {
		return err
	}

	return nil
}
