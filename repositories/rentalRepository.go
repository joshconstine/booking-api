package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type RentalRepository interface {
	FindAll() []response.RentalResponse
	FindById(id uint) response.RentalInformationResponse
	Create(rental request.CreateRentalRequest) (response.RentalResponse, error)
}
