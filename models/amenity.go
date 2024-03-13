package models

import (
	"gorm.io/gorm"
)

type Amenity struct {
	gorm.Model
	Name          string
	AmenityTypeID uint
}

func (a *Amenity) TableName() string {
	return "amenities"
}
