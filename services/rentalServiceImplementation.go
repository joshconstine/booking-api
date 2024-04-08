package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type RentalServiceImplementation struct {
	RentalRepository repositories.RentalRepository
	Validate         *validator.Validate
}

func NewRentalServiceImplementation(rentalRepository repositories.RentalRepository, validate *validator.Validate) RentalService {
	return &RentalServiceImplementation{
		RentalRepository: rentalRepository,
		Validate:         validate,
	}
}

func (t *RentalServiceImplementation) FindAll() []response.RentalResponse {
	return t.RentalRepository.FindAll()
}

func (t *RentalServiceImplementation) FindById(id uint) response.RentalInformationResponse {
	return t.RentalRepository.FindById(id)
}
