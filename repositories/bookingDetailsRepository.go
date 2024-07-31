package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
)

type BookingDetailsRepository interface {
	FindById(id uint) response.BookingDetailsResponse
	FindByBookingId(id string) response.BookingDetailsResponse
	Create(details models.BookingDetails) response.BookingDetailsResponse
	Update(details request.UpdateBookingDetailsRequest) (response.BookingDetailsResponse, error)
	UpdatePaymentCompleteStatus(bookingId string, status bool)
}
