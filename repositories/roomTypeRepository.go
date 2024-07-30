package repositories

import (
	"booking-api/data/response"
)

type RoomTypeRepository interface {
	FindAll() []response.RoomTypeResponse
	FindById(id uint) response.RoomTypeResponse
	Create(rentalRoomTypeName string) response.RoomTypeResponse
}
