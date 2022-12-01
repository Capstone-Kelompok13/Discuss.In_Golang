package models

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func (Topic) TableName() string {
	return "Topics"
}
