package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type RentalService interface {
	FindAll() []response.RentalResponse
	FindById(id uint) response.RentalInformationResponse
	Create(rental request.CreateRentalRequest) (response.RentalResponse, error)
	CreateStep1(rental request.CreateRentalStep1Params) (response.RentalResponse, error)
	Update(rental request.UpdateRentalRequest) (response.RentalResponse, error)
	UpdateRental(rental request.CreateRentalStep1Params) (response.RentalResponse, error)
}
