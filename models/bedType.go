package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BedType struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (b *BedType) TableName() string {
	return "bed_types"
}

func (b *BedType) MapBedTypeToResponse() response.BedTypeResponse {
	return response.BedTypeResponse{
		ID:   b.ID,
		Name: b.Name,
	}
}
