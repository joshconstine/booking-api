package services

import (
	responses "booking-api/data/response"
)

type BookingCostTypeService interface {
	FindAll() []responses.BookingCostTypeResponse
	FindById(id uint) responses.BookingCostTypeResponse
}
