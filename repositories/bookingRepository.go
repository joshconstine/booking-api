package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
)

type BookingRepository interface {
	FindAll() []response.BookingResponse
	FindById(id string) response.BookingInformationResponse
	Create(booking *request.CreateBookingRequest) (string, error)
	Update(booking models.Booking) models.Booking
	CheckIfEntitiesCanBeBooked(request *request.CreateBookingRequest) (bool, error)
}
