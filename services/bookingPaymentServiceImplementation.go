package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BookingPaymentServiceImplementation struct {
	bookingPaymentRepository repositories.BookingPaymentRepository
	Validate                 *validator.Validate
}

func NewBookingPaymentServiceImplementation(bookingPaymentRepository repositories.BookingPaymentRepository, validate *validator.Validate) BookingPaymentService {
	return &BookingPaymentServiceImplementation{
		bookingPaymentRepository: bookingPaymentRepository,
		Validate:                 validate,
	}
}

func (t BookingPaymentServiceImplementation) Create(bookingPayment requests.CreateBookingPaymentRequest) (response.BookingPaymentResponse, error) {
	err := t.Validate.Struct(bookingPayment)

	if err != nil {
		return response.BookingPaymentResponse{}, err
	}

	return t.bookingPaymentRepository.Create(bookingPayment)

}

func (t BookingPaymentServiceImplementation) FindAll() []response.BookingPaymentResponse {
	result := t.bookingPaymentRepository.FindAll()

	return result
}

func (t BookingPaymentServiceImplementation) FindById(id uint) response.BookingPaymentResponse {
	result := t.bookingPaymentRepository.FindById(id)

	return result
}

func (t BookingPaymentServiceImplementation) FindByBookingId(id string) []response.BookingPaymentResponse {
	result := t.bookingPaymentRepository.FindByBookingId(id)

	return result
}

func (t BookingPaymentServiceImplementation) FindTotalAmountByBookingId(id string) float64 {
	result := t.bookingPaymentRepository.FindTotalAmountByBookingId(id)

	return result
}
