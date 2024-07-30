package services

import (
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/repositories"
)

type BookingDetailsServiceImplementation struct {
	bookingDetailsRepository repositories.BookingDetailsRepository
}

func NewBookingDetailsServiceImplementation(bookingDetailsRepository repositories.BookingDetailsRepository) BookingDetailsServiceImplementation {
	return BookingDetailsServiceImplementation{bookingDetailsRepository: bookingDetailsRepository}
}

func (service BookingDetailsServiceImplementation) FindById(id uint) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.FindById(id)
}

func (service BookingDetailsServiceImplementation) Create(details models.BookingDetails) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.Create(details)
}

func (service BookingDetailsServiceImplementation) Update(details models.BookingDetails) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.Update(details)
}
