package repositories

import (
	"booking-api/models"
)

type BookingRepository interface {
	FindAll() []models.Booking
	FindById(id string) models.Booking
	Create(booking models.Booking) models.Booking
	Update(booking models.Booking) models.Booking
}
