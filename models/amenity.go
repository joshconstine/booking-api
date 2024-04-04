package models

import (
	"gorm.io/gorm"
)

type Amenity struct {
	gorm.Model
	Name          string
	AmenityTypeID uint
	AmenityType   AmenityType
}

func (a *Amenity) TableName() string {
	return "amenities"
}
