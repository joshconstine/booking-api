package models

import (
	"gorm.io/gorm"
)

type EntityPhoto struct {
	gorm.Model
	PhotoID    uint
	EntityID   uint   `gorm:"primaryKey"`
	EntityType string `gorm:"primaryKey"`
	Photo      Photo
}

func (e *EntityPhoto) TableName() string {
	return "entity_photos"
}
