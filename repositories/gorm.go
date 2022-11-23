package repositories

import (
	"discusiin/models"

	"gorm.io/gorm"
)

type GormSql struct {
	DB *gorm.DB
}

// // SaveTopic implements IDatabase
// func (*GormSql) SaveTopic(models.Topic) error {
// 	panic("unimplemented")
// }

func NewGorm(db *gorm.DB) IDatabase {
	return &GormSql{
		DB: db,
	}
}

// User
func (db GormSql) SaveNewUser(user models.User) error {
	result := db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (db GormSql) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.DB.
		Where("username = ?",
			username).
		First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (db GormSql) Login(email, password string) (models.User, error) {
	var user models.User
	err := db.DB.
		Where("email = ? AND password = ?",
			email, password).
		First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

//Topic
func (db GormSql) GetTopicByName(name string) (models.Topic, error) {
	var topic models.Topic
	err := db.DB.Where("name = ?", name).First(&topic).Error

	if err != nil {
		return models.Topic{}, err
	}

	return topic, nil
}

func (db GormSql) GetTopicByID(id int) (models.Topic, error) {
	var topic models.Topic
	err := db.DB.Where("id = ?", id).First(&topic).Error

	if err != nil {
		return models.Topic{}, err
	}

	return topic, nil
}

func (db GormSql) SaveNewTopic(topic models.Topic) error {
	result := db.DB.Create(&topic)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db GormSql) SaveTopic(topic models.Topic) error {
	err := db.DB.Where("id = ?", topic.ID).Save(topic)
	if err != nil {
		return err.Error
	}
	return nil
}

func (db GormSql) SaveNewModerator(userId int, topicId uint) error {
	var mod models.Moderator

	mod.UserID = userId
	mod.TopicID = int(topicId)

	result := db.DB.Create(&mod)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
