package repositories

import (
	responses "booking-api/data/response"
)

type RentalStatusRepository interface {
	FindAll() []responses.RentalStatusResponse
	FindByRentalId(rentalId uint) responses.RentalStatusResponse
	UpdateStatusForRentalId(rentalId uint, isClean bool) responses.RentalStatusResponse
}
