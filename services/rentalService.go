package services

import (
	"booking-api/models"
)

type RentalService interface {
	FindAll() []models.Rental
	FindById(id uint) models.Rental
	// Create(rental models.Rental) models.Rental
}
