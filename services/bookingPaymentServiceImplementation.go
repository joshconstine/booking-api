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

func (t BookingPaymentServiceImplementation) Create(bookingPayment requests.CreateBookingPaymentRequest) response.BookingPaymentResponse {
	err := t.Validate.Struct(bookingPayment)

	if err != nil {
		panic(err)
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
