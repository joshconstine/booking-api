package repositories

import (
	"booking-api/models"
)

type BookingDetailsRepository interface {
	FindById(id uint) models.BookingDetails
	Create(details models.BookingDetails) models.BookingDetails
	Update(details models.BookingDetails) models.BookingDetails
}
