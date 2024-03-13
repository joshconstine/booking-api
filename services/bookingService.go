package services

import responses "booking-api/data/response"

type BookingService interface {
	FindAll() []responses.BookingResponse
	FindById(id uint) responses.BookingResponse
}
