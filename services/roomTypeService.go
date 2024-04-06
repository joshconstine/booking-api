package services

import "booking-api/data/response"

type RoomTypeService interface {
	FindAll() []response.RoomTypeResponse
	FindById(id int) response.RoomTypeResponse
	Create(name string) response.RoomTypeResponse
}
