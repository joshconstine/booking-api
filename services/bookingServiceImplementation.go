package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BookingServiceImplementation struct {
	BookingRepository repositories.BookingRepository
	Validate          *validator.Validate
}

func NewBookingServiceImplementation(bookingRepository repositories.BookingRepository, validate *validator.Validate) BookingService {
	return &BookingServiceImplementation{
		BookingRepository: bookingRepository,
		Validate:          validate,
	}
}

func (t BookingServiceImplementation) FindAll() []response.BookingResponse {
	result := t.BookingRepository.FindAll()

	var bookings []response.BookingResponse
	for _, value := range result {
		booking := response.BookingResponse{
			ID:               value.ID,
			BookingStatusID:  value.BookingStatusID,
			BookingDetailsID: value.BookingDetailsID,
		}
		bookings = append(bookings, booking)
	}
	return bookings
}

func (t BookingServiceImplementation) FindById(id uint) response.BookingResponse {
	result := t.BookingRepository.FindById(id)

	booking := response.BookingResponse{
		ID:               result.ID,
		BookingStatusID:  result.BookingStatusID,
		BookingDetailsID: result.BookingDetailsID,
	}
	return booking
}
