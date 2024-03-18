package repositories

import (
	"booking-api/models"
)

type AmenityRepository interface {
	FindAll() []models.Amenity
	FindById(id uint) models.Amenity
	Create(amenity models.Amenity) models.Amenity
}
