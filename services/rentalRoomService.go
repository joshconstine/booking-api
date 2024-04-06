package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type RentalRoomService interface {
	FindAll() []response.RentalRoomResponse
	FindById(id uint) response.RentalRoomResponse
	Create(rentalRoom request.RentalRoomCreateRequest) (response.RentalRoomResponse, error)
}
