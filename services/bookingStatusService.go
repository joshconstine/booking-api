package services

import (
	responses "booking-api/data/response"
)

type BookingStatusService interface {
	FindAll() []responses.BookingStatusResponse
	FindById(id uint) responses.BookingStatusResponse
}
