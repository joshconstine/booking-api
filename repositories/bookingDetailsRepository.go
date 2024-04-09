package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type BookingDetailsRepository interface {
	FindById(id uint) response.BookingDetailsResponse
	Create(details models.BookingDetails) response.BookingDetailsResponse
	Update(details models.BookingDetails) response.BookingDetailsResponse
}
