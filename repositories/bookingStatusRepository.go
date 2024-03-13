package repositories

import (
	"booking-api/models"
)

type BookingStatusRepository interface {
	FindAll() []models.BookingStatus
	FindById(id uint) models.BookingStatus
}
