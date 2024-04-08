package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type RentalRepository interface {
	FindAll() []response.RentalResponse
	FindById(id uint) response.RentalInformationResponse
	Create(rental models.Rental) models.Rental
}
