package services

import (
	responses "booking-api/data/response"
	"booking-api/models"
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

func (t *RentalServiceImplementation) FindAll() []models.Rental {
	return t.RentalRepository.FindAll()
}

func (t *RentalServiceImplementation) FindById(id uint) responses.RentalInformationResponse {
	return t.RentalRepository.GetInformationForRental(id)
}
