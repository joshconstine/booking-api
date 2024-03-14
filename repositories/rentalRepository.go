package repositories

import (
	"booking-api/models"
)

type RentalRepository interface {
	FindAll() []models.Rental
	FindById(id uint) models.Rental
	Create(rental models.Rental) models.Rental
}
