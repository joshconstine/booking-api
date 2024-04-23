package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
)

type BookingRepository interface {
	FindAll() []response.BookingResponse
	FindById(id string) response.BookingInformationResponse
	GetSnapshot() []response.BookingSnapshotResponse
	Create(booking *request.CreateBookingRequest) (string, error)
	Update(booking models.Booking) models.Booking
	CheckIfEntitiesCanBeBooked(request *request.CreateBookingRequest) (bool, error)
}
