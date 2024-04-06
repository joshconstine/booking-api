package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type RentalRoomRepository interface {
	FindAll() []response.RentalRoomResponse
	FindById(id uint) response.RentalRoomResponse
	Create(rentalRoom request.RentalRoomCreateRequest) response.RentalRoomResponse
}
