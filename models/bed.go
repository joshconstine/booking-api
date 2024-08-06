package models

import (
	"booking-api/data/response"
	"gorm.io/gorm"
)

type Bed struct {
	gorm.Model
	RentalRoomID uint `gorm:"not null"`
	BedTypeID    uint `gorm:"not null"`
	BedType      BedType
}

func (b *Bed) MapBedToResponse() response.BedResponse {
	return response.BedResponse{
		ID:   b.ID,
		Name: b.BedType.Name,
	}
}
