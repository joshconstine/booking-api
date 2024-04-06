package services

import (
	"booking-api/data/response"
	"booking-api/repositories"
)

type RoomTypeServiceImplementation struct {
	RoomTypeRepository repositories.RoomTypeRepository
}

func NewRoomTypeServiceImplementation(roomTypeRepository repositories.RoomTypeRepository) RoomTypeService {
	return &RoomTypeServiceImplementation{
		RoomTypeRepository: roomTypeRepository,
	}
}

func (r *RoomTypeServiceImplementation) FindAll() []response.RoomTypeResponse {
	return r.RoomTypeRepository.FindAll()
}

func (r *RoomTypeServiceImplementation) FindById(id int) response.RoomTypeResponse {
	return r.RoomTypeRepository.FindById(uint(id))
}

func (r *RoomTypeServiceImplementation) Create(name string) response.RoomTypeResponse {
	return r.RoomTypeRepository.Create(name)
}
