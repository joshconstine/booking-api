package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	rentals "booking-api/view/rentals"
)

type RentalService interface {
	FindAll() []response.RentalResponse
	FindById(id uint) response.RentalInformationResponse
	Create(rental request.CreateRentalRequest) (response.RentalResponse, error)
	Update(rental request.UpdateRentalRequest) (response.RentalResponse, error)
	UpdateRental(rental rentals.RentalFormParams) (response.RentalResponse, error)
}
