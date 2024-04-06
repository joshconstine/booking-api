package repositories

import (
	"booking-api/data/response"
)

type RoomTypeRepository interface {
	FindAll() []response.RentalRoomTypeResponse
	FindById(id uint) response.RentalRoomTypeResponse
	Create(rentalRoomTypeName string) response.RentalRoomTypeResponse
}
