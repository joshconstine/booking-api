package services

import (
	"booking-api/data/request"
	responses "booking-api/data/response"
	"booking-api/models"
)

type BookingDetailsService interface {
	FindById(id uint) responses.BookingDetailsResponse
	FindByBookingId(id string) responses.BookingDetailsResponse
	Create(details models.BookingDetails) responses.BookingDetailsResponse
	Update(details request.UpdateBookingDetailsRequest) (responses.BookingDetailsResponse, error)
	AuditBookingDetailsForBooking(bookingID string) error
}
