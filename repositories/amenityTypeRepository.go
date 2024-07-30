package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type AmenityTypeRepository interface {
	FindAll() []response.AmenityTypeResponse
	FindById(id uint) response.AmenityTypeResponse
	Create(amenityType models.AmenityType) response.AmenityTypeResponse
}
