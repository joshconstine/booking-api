package services

import (
	requests "booking-api/data/request"
	responses "booking-api/data/response"
)

type BookingService interface {
	FindAll() []responses.BookingResponse
	FindById(id uint) responses.BookingResponse
	Create(request requests.CreateUserRequest) responses.BookingResponse
}