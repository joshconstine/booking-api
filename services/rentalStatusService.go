package services

import (
	"booking-api/data/response"
)

type RentalStatusService interface {
	FindAll() []response.RentalStatusResponse
	FindByRentalId(rentalId uint) response.RentalStatusResponse
	UpdateStatusForRentalId(rentalId uint, isClean bool) response.RentalStatusResponse
}
