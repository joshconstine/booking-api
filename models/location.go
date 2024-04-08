package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (l *Location) TableName() string {
	return "locations"
}

func (l *Location) MapLocationToResponse() response.LocationResponse {
	return response.LocationResponse{
		ID:   l.ID,
		Name: l.Name,
	}
}
