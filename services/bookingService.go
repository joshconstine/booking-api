package services

import (
	"booking-api/data/request"
	responses "booking-api/data/response"
)

type BookingService interface {
	Create(request *request.CreateBookingRequest) (string, error)
	FindAll() []responses.BookingResponse
	FindById(id string) responses.BookingInformationResponse
}
