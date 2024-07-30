package services

import (
	responses "booking-api/data/response"
	"booking-api/models"
)

type BookingDetailsService interface {
	FindById(id uint) responses.BookingDetailsResponse
	Create(details models.BookingDetails) responses.BookingDetailsResponse
	Update(details models.BookingDetails) responses.BookingDetailsResponse
}
