package services

import (
	"booking-api/data/response"
	"booking-api/repositories"
)

type BookingDetailsServiceImplementation struct {
	bookingDetailsRepository repositories.BookingDetailsRepository
}

func NewBookingDetailsServiceImplementation(bookingDetailsRepository repositories.BookingDetailsRepository) BookingDetailsServiceImplementation {
	return BookingDetailsServiceImplementation{bookingDetailsRepository: bookingDetailsRepository}
}

func (service BookingDetailsServiceImplementation) FindById(id uint) response.BookingDetailsResponse {
	bookingDetails := service.bookingDetailsRepository.FindById(id)
	return response.BookingDetailsResponse{
		ID:               bookingDetails.ID,
		BookingID:        uint(bookingDetails.BookingID),
		PaymentComplete:  bookingDetails.PaymentComplete,
		PaymentDueDate:   bookingDetails.PaymentDueDate,
		DocumentsSigned:  bookingDetails.DocumentsSigned,
		BookingStartDate: bookingDetails.BookingStartDate,
		InvoiceID:        bookingDetails.InvoiceID,
	}
}
