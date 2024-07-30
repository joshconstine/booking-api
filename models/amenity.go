package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Amenity struct {
	gorm.Model
	Name          string `gorm:"index:idx_entity,unique; not null"`
	AmenityTypeID uint   `gorm:"index:idx_entity,unique; not null"`
	AmenityType   AmenityType
}

func (a *Amenity) TableName() string {
	return "amenities"
}

func (a *Amenity) MapAmenityToResponse() response.AmenityResponse {
	response := response.AmenityResponse{
		ID:   a.ID,
		Name: a.Name,
	}

	response.AmenityType = a.AmenityType.MapAmenityTypeToResponse()

	return response
}
