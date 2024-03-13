package repositories

import (
	"booking-api/models"
)

type BookingDetailsRepository interface {
	FindById(id uint) models.BookingDetails
}
