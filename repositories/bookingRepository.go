package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type BookingRepository interface {
	FindAll() []response.BookingResponse
	FindById(id string) response.BookingInformationResponse
	Create(booking models.Booking) models.Booking
	Update(booking models.Booking) models.Booking
}
