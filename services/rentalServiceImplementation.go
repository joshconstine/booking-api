package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
	rentals "booking-api/view/rentals"

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

func (t *RentalServiceImplementation) Create(rental request.CreateRentalRequest) (response.RentalResponse, error) {
	err := t.Validate.Struct(rental)

	if err != nil {
		panic(err)
	}

	return t.RentalRepository.Create(rental)
}

func (t *RentalServiceImplementation) Update(rental request.UpdateRentalRequest) (response.RentalResponse, error) {
	err := t.Validate.Struct(rental)

	if err != nil {
		panic(err)
	}

	return t.RentalRepository.Update(rental)
}

func (t *RentalServiceImplementation) UpdateRental(rental rentals.RentalFormParams) (response.RentalResponse, error) {
	err := t.Validate.Struct(rental)

	if err != nil {
		panic(err)
	}

	return t.RentalRepository.UpdateRental(rental)
}
