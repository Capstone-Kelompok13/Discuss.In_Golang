package models

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Suspended   time.Time `json:"suspended" form:"suspended"`
}

func (Topic) TableName() string {
	return "topics"
}
