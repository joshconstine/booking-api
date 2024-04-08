package services

import (
	"booking-api/data/response"
)

type RentalService interface {
	FindAll() []response.RentalResponse
	FindById(id uint) response.RentalInformationResponse
	// Create(rental models.Rental) models.Rental
}
