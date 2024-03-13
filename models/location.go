package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name string
}

func (l *Location) TableName() string {

	return "locations"
}
