package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type RentalRoomRepository interface {
	FindAll() []response.RentalRoomResponse
	FindById(id uint) response.RentalRoomResponse
	FindByRentalId(rentalId uint) []response.RentalRoomResponse
	AddBedToRoom(roomId uint, bedId uint) error
	Create(rentalRoom request.RentalRoomCreateRequest) response.RentalRoomResponse
	Update(rentalRoom request.UpdateRentalRoomRequest) response.RentalRoomResponse
	Delete(id uint) error
}
