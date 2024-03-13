package services

import (
	responses "booking-api/data/response"
)

type BookingDetailsService interface {
	FindById(id uint) responses.BookingDetailsResponse
}
