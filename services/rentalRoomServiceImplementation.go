package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type RentalRoomServiceImplementation struct {
	RentalRoomRepository repositories.RentalRoomRepository
	Validate             *validator.Validate
}

func NewRentalRoomServiceImplementation(rentalRoomRepository repositories.RentalRoomRepository, validate *validator.Validate) RentalRoomService {

	return &RentalRoomServiceImplementation{
		RentalRoomRepository: rentalRoomRepository,
		Validate:             validate,
	}
}

func (r *RentalRoomServiceImplementation) FindAll() []response.RentalRoomResponse {
	return r.RentalRoomRepository.FindAll()
}

func (r *RentalRoomServiceImplementation) FindById(id uint) response.RentalRoomResponse {
	return r.RentalRoomRepository.FindById(id)
}

func (r *RentalRoomServiceImplementation) FindByRentalId(rentalId uint) []response.RentalRoomResponse {
	return r.RentalRoomRepository.FindByRentalId(rentalId)
}

func (r *RentalRoomServiceImplementation) Create(rentalRoom request.RentalRoomCreateRequest) (response.RentalRoomResponse, error) {

	err := r.Validate.Struct(rentalRoom)

	if err != nil {
		return response.RentalRoomResponse{}, err
	}

	return r.RentalRoomRepository.Create(rentalRoom), nil

}

func (r *RentalRoomServiceImplementation) Update(rentalRoom request.UpdateRentalRoomRequest) (response.RentalRoomResponse, error) {

	err := r.Validate.Struct(rentalRoom)

	if err != nil {
		return response.RentalRoomResponse{}, err
	}

	return r.RentalRoomRepository.Update(rentalRoom), nil

}

func (r *RentalRoomServiceImplementation) Delete(id uint) error {
	return r.RentalRoomRepository.Delete(id)

}

func (r *RentalRoomServiceImplementation) DeleteBed(id uint) error {
	return r.RentalRoomRepository.DeleteBed(id)
}

func (r *RentalRoomServiceImplementation) AddBedToRoom(roomId uint, bedId uint) error {
	return r.RentalRoomRepository.AddBedToRoom(roomId, bedId)
}
