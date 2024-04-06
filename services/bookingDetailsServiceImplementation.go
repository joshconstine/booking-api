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
	bookingDetails := service.bookingDetailsRepository.FindById(id)
	return response.BookingDetailsResponse{
		ID:               bookingDetails.ID,
		BookingID:        bookingDetails.BookingID,
		PaymentComplete:  bookingDetails.PaymentComplete,
		PaymentDueDate:   bookingDetails.PaymentDueDate,
		DocumentsSigned:  bookingDetails.DocumentsSigned,
		BookingStartDate: bookingDetails.BookingStartDate,
		InvoiceID:        bookingDetails.InvoiceID,
	}
}

func (service BookingDetailsServiceImplementation) Create(details models.BookingDetails) response.BookingDetailsResponse {
	bookingDetails := service.bookingDetailsRepository.Create(details)
	return response.BookingDetailsResponse{
		ID:               bookingDetails.ID,
		BookingID:        bookingDetails.BookingID,
		PaymentComplete:  bookingDetails.PaymentComplete,
		PaymentDueDate:   bookingDetails.PaymentDueDate,
		DocumentsSigned:  bookingDetails.DocumentsSigned,
		BookingStartDate: bookingDetails.BookingStartDate,
		InvoiceID:        bookingDetails.InvoiceID,
	}
}

func (service BookingDetailsServiceImplementation) Update(details models.BookingDetails) response.BookingDetailsResponse {

	bookingDetails := service.bookingDetailsRepository.Update(details)
	return response.BookingDetailsResponse{
		ID:               bookingDetails.ID,
		BookingID:        bookingDetails.BookingID,
		PaymentComplete:  bookingDetails.PaymentComplete,
		PaymentDueDate:   bookingDetails.PaymentDueDate,
		DocumentsSigned:  bookingDetails.DocumentsSigned,
		BookingStartDate: bookingDetails.BookingStartDate,
		InvoiceID:        bookingDetails.InvoiceID,
	}
}
