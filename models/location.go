package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (l *Location) TableName() string {
	return "locations"
}
