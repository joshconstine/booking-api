package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type AmenityType struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (a *AmenityType) TableName() string {
	return "amenity_types"
}

func (a *AmenityType) MapAmenityTypeToResponse() response.AmenityTypeResponse {
	return response.AmenityTypeResponse{
		ID:   a.ID,
		Name: a.Name,
	}
}
