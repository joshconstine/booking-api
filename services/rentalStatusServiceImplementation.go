package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type RentalStatusServiceImplementation struct {
	RentalStatusRepository repositories.RentalStatusRepository
	Validate               *validator.Validate
}

func NewRentalStatusServiceImplementation(rentalStatusRepository repositories.RentalStatusRepository, validate *validator.Validate) RentalStatusService {
	return &RentalStatusServiceImplementation{
		RentalStatusRepository: rentalStatusRepository,
		Validate:               validate,
	}
}

func (t *RentalStatusServiceImplementation) FindAll() []response.RentalStatusResponse {
	return t.RentalStatusRepository.FindAll()
}

func (t *RentalStatusServiceImplementation) FindByRentalId(rentalId uint) response.RentalStatusResponse {
	return t.RentalStatusRepository.FindByRentalId(rentalId)
}

func (t *RentalStatusServiceImplementation) UpdateStatusForRentalId(rentalId uint, isClean bool) response.RentalStatusResponse {
	return t.RentalStatusRepository.UpdateStatusForRentalId(rentalId, isClean)
}
