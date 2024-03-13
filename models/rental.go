package models

import (
	"gorm.io/gorm"
)

type Rental struct {
	gorm.Model
	Name        string
	LocationID  uint
	Bedrooms    int
	Bathrooms   int
	Description string
	Location    Location
	Amenities   []Amenity `gorm:"many2many:rental_amenities;"`
}

func (r *Rental) TableName() string {
	return "rentals"
}
