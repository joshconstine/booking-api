package models

import (
	"gorm.io/gorm"
)

type AmenityType struct {
	gorm.Model
	Name string
}

func (a *AmenityType) TableName() string {
	return "amenity_types"
}
