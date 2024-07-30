package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityPhoto struct {
	gorm.Model
	EntityID   uint   `gorm:"primaryKey"`
	EntityType string `gorm:"primaryKey"`
	PhotoID    uint   `gorm:"not null"`
	Photo      Photo
}

func (e *EntityPhoto) TableName() string {
	return "entity_photos"
}

func (e *EntityPhoto) MapEntityPhotoToResponse() response.EntityPhotoResponse {
	response := response.EntityPhotoResponse{
		ID:    e.ID,
		Photo: e.Photo.MapPhotoToResponse(),
	}

	return response
}
