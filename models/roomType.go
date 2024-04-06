package models

import (
	"gorm.io/gorm"
)

type RoomType struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
