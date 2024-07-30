package services

import (
	"booking-api/data/request"
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
func (service BookingDetailsServiceImplementation) FindByBookingId(id string) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.FindByBookingId(id)
}

func (service BookingDetailsServiceImplementation) Create(details models.BookingDetails) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.Create(details)
}

func (service BookingDetailsServiceImplementation) Update(details request.UpdateBookingDetailsRequest) (response.BookingDetailsResponse, error) {
	return service.bookingDetailsRepository.Update(details)
}
