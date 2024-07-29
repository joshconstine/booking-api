package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type BookingPaymentService interface {
	FindAll() []response.BookingPaymentResponse
	FindById(id uint) response.BookingPaymentResponse
	FindByBookingId(id string) []response.BookingPaymentResponse
	FindTotalAmountByBookingId(id string) float64
	Create(bookingPayment requests.CreateBookingPaymentRequest) (response.BookingPaymentResponse, error)
}
