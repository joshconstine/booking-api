package services

import (
	"booking-api/constants"
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
	"errors"

	"github.com/go-playground/validator/v10"
)

type BookingPaymentServiceImplementation struct {
	bookingPaymentRepository repositories.BookingPaymentRepository
	BookingCostItemService   BookingCostItemService
	BookingDetailsService    BookingDetailsService
	Validate                 *validator.Validate
}

func NewBookingPaymentServiceImplementation(bookingPaymentRepository repositories.BookingPaymentRepository, bookingDetailsService BookingDetailsService, bookingCostItemService BookingCostItemService, validate *validator.Validate) *BookingPaymentServiceImplementation {
	return &BookingPaymentServiceImplementation{bookingPaymentRepository: bookingPaymentRepository, BookingDetailsService: bookingDetailsService, BookingCostItemService: bookingCostItemService, Validate: validate}
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

	// Check if payment is completed
	if t.CheckIfPaymentIsCompleted(bookingPayment.BookingID) {
		return response.BookingPaymentResponse{}, errors.New(constants.PAYMENT_IS_ALREADY_COMPLETED_ERROR)

	}

	response, err := t.bookingPaymentRepository.Create(bookingPayment)

	if err != nil {
		return response, err

	}

	//check if payment is completed
	if t.bookingPaymentRepository.FindTotalOutstandingAmountByBookingId(response.BookingID) == 0 {
		//update booking status to paid
		bookingDetails := t.BookingDetailsService.FindByBookingId(response.BookingID)

		var request requests.UpdateBookingDetailsRequest
		request.ID = bookingDetails.ID
		request.PaymentComplete = true
		request.BookingStartDate = bookingDetails.BookingStartDate
		request.PaymentDueDate = bookingDetails.PaymentDueDate
		request.DocumentsSigned = bookingDetails.DocumentsSigned
		request.DepositPaid = bookingDetails.DepositPaid
		request.GuestCount = bookingDetails.GuestCount

		if bookingDetails.PaymentComplete == false {
			bookingDetails.PaymentComplete = true

			_, err := t.BookingDetailsService.Update(request)

			if err != nil {
				return response, err

			}

		}
	}
	return response, nil
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
