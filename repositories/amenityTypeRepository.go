package repositories

import (
	"booking-api/models"
)

type AmenityTypeRepository interface {
	FindAll() []models.AmenityType
	FindById(id uint) models.AmenityType
	Create(amenityType models.AmenityType) models.AmenityType
}
