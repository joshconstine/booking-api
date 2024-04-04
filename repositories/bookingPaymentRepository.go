package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type BookingPaymentRepository interface {
	FindAll() []response.BookingPaymentResponse
	FindById(id uint) response.BookingPaymentResponse
	FindByBookingId(id uint) []response.BookingPaymentResponse
	FindTotalAmountByBookingId(id uint) float64
	Create(bookingPayment requests.CreateBookingPaymentRequest) response.BookingPaymentResponse
}
