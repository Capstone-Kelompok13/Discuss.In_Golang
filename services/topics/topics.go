package topics

import (
	"discusiin/models"
	"discusiin/repositories"
	"errors"
)

func NewTopicServices(db repositories.IDatabase) ITopicServices {
	return &topicServices{IDatabase: db}
}

type ITopicServices interface {
	CreateTopic(topic models.Topic) error
	GetTopic(id int) (models.Topic, error)
	SaveTopic(topic models.Topic) error
}

type topicServices struct {
	repositories.IDatabase
}

func (t *topicServices) CreateTopic(topic models.Topic) error {
	//check if topic already exist
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

func (t *topicServices) SaveTopic(topic models.Topic) error {
	err := t.IDatabase.SaveTopic(topic)
	if err != nil {
		return err
	}

	return nil
}
