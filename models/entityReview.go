package models

import (
	"gorm.io/gorm"
)

type EntityReview struct {
	gorm.Model
	EntityID   uint    `gorm:"not null"`
	EntityType string  `gorm:"not null"`
	Rating     float64 `gorm:"not null"`
	Message    string
}

func (c *EntityReview) TableName() string {
	return "entity_reviews"
}
