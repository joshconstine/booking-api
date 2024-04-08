package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type RentalRepository interface {
	FindAll() []response.RentalResponse
	FindById(id uint) models.Rental
	Create(rental models.Rental) models.Rental
}
