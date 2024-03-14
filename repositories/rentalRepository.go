package repositories

import (
	responses "booking-api/data/response"
	"booking-api/models"
)

type RentalRepository interface {
	FindAll() []models.Rental
	FindById(id uint) models.Rental
	GetInformationForRental(id uint) responses.RentalInformationResponse
	Create(rental models.Rental) models.Rental
}
