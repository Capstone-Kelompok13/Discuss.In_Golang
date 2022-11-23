package topics

import (
	"discusiin/models"
	"discusiin/repositories"
	"errors"
)

func NewTopicServices(db repositories.ITopicDatabase) ITopicServices {
	return &topicServices{ITopicDatabase: db}
}

type ITopicServices interface {
	CreateTopic(topic models.Topic) error
}

type topicServices struct {
	repositories.ITopicDatabase
}

func (t *topicServices) CreateTopic(topic models.Topic) error {
	//check if topic already exist
	_, err1 := t.ITopicDatabase.GetTopicByName(topic.Name)
	if err1 != nil {
		if err1.Error() == "record not found" {
			err2 := t.ITopicDatabase.SaveNewTopic(topic)
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

	//add creator as moderator if topic created for the first time
	topicData, err3 := t.ITopicDatabase.GetTopicByName(topic.Name)
	if err3 != nil {
		return err3
	} else {
		err4 := t.ITopicDatabase.SaveNewModerator(topicData.UserID, topicData.ID)
		if err4 != nil {
			return err4
		}
	}

	return nil
}
