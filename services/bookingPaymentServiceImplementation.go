package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BookingPaymentServiceImplementation struct {
	bookingPaymentRepository repositories.BookingPaymentRepository
	BookingCostItemService   BookingCostItemService
	Validate                 *validator.Validate
}

func NewBookingPaymentServiceImplementation(bookingPaymentRepository repositories.BookingPaymentRepository, bookingCostItemService BookingCostItemService, validate *validator.Validate) BookingPaymentService {
	return BookingPaymentServiceImplementation{bookingPaymentRepository: bookingPaymentRepository, BookingCostItemService: bookingCostItemService, Validate: validate}
}

func (t BookingPaymentServiceImplementation) CheckIfPaymentIsCompleted(bookingId string) bool {
	result := t.bookingPaymentRepository.CheckIfPaymentIsCompleted(bookingId)

	return result
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

func (t BookingPaymentServiceImplementation) FindTotalPaidByBookingId(id string) float64 {
	result := t.bookingPaymentRepository.FindTotalPaidByBookingId(id)

	return result
}
