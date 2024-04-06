package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type BookingPaymentRepository interface {
	FindAll() []response.BookingPaymentResponse
	FindById(id uint) response.BookingPaymentResponse
	FindByBookingId(id string) []response.BookingPaymentResponse
	FindTotalAmountByBookingId(id string) float64
	Create(bookingPayment requests.CreateBookingPaymentRequest) response.BookingPaymentResponse
}
