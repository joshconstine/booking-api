package services

import (
	"booking-api/data/request"
	responses "booking-api/data/response"
)

type BookingService interface {
	Create(request *request.CreateBookingRequest) (string, error)
	CreateBookingWithUserInformation(request *request.CreateBookingWithUserInformationRequest) (string, error)
	FindAll() []responses.BookingResponse
	GetSnapshot(request request.GetBookingSnapshotRequest) []responses.BookingSnapshotResponse
	FindById(id string) (responses.BookingInformationResponse, error)
}
