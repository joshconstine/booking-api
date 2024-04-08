package models

import (
	"booking-api/data/response"

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

func (e *EntityPhoto) MapEntityPhotoToResponse() response.PhotoResponse {
	return response.PhotoResponse{
		ID:  e.ID,
		URL: e.Photo.URL,
	}
}
