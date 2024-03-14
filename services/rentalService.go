package services

import (
	"booking-api/data/response"
	"booking-api/models"
)

type RentalService interface {
	FindAll() []models.Rental
	FindById(id uint) response.RentalInformationResponse
	// Create(rental models.Rental) models.Rental
}
