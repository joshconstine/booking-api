package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	rentals "booking-api/view/rentals"
)

type RentalRepository interface {
	FindAll() []response.RentalResponse
	FindAllIDs() []uint
	FindById(id uint) response.RentalInformationResponse
	Create(rental request.CreateRentalRequest) (response.RentalResponse, error)
	Update(rental request.UpdateRentalRequest) (response.RentalResponse, error)
	UpdateRental(rental rentals.RentalFormParams) (response.RentalResponse, error)
}
