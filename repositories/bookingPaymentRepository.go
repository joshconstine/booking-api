package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type BookingPaymentRepository interface {
	FindAll() []response.BookingPaymentResponse
	FindById(id uint) response.BookingPaymentResponse
	FindByBookingId(id string) []response.BookingPaymentResponse
	FindTotalPaidByBookingId(id string) float64
	CheckIfPaymentIsCompleted(id string) bool
	Create(bookingPayment requests.CreateBookingPaymentRequest) (response.BookingPaymentResponse, error)
}
