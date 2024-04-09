package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	responses "booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type BookingPaymentRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingPaymentRepositoryImplementation(db *gorm.DB) BookingPaymentRepository {
	return &BookingPaymentRepositoryImplementation{Db: db}
}

func (t *BookingPaymentRepositoryImplementation) FindAll() []response.BookingPaymentResponse {
	var bookingPayments []models.BookingPayment
	var response []response.BookingPaymentResponse

	result := t.Db.Preload("PaymentMethod").Find(&bookingPayments)
	if result.Error != nil {
		return []responses.BookingPaymentResponse{}
	}

	for _, bookingPayment := range bookingPayments {

		response = append(response, bookingPayment.MapBookingPaymentToResponse())

	}
	return response
}

func (t *BookingPaymentRepositoryImplementation) FindById(id uint) response.BookingPaymentResponse {
	var bookingPayment models.BookingPayment

	result := t.Db.Model(&models.BookingPayment{}).Preload("PaymentMethod").First(&bookingPayment, id)
	if result.Error != nil {

		return response.BookingPaymentResponse{}
	}

	return bookingPayment.MapBookingPaymentToResponse()
}

func (t *BookingPaymentRepositoryImplementation) Create(bookingPayment requests.CreateBookingPaymentRequest) response.BookingPaymentResponse {
	bookingPaymentModel := models.BookingPayment{
		BookingID:       bookingPayment.BookingID,
		PaymentMethodID: bookingPayment.PaymentMethodID,
		PaymentAmount:   bookingPayment.PaymentAmount,
	}
	result := t.Db.Create(&bookingPaymentModel)
	if result.Error != nil {
		return response.BookingPaymentResponse{}
	}

	return bookingPaymentModel.MapBookingPaymentToResponse()
}
func (t *BookingPaymentRepositoryImplementation) FindByBookingId(id string) []response.BookingPaymentResponse {
	var bookingPayments []models.BookingPayment

	result := t.Db.Model(&models.BookingPayment{}).Where("booking_id = ?", id).Preload("PaymentMethod").Find(&bookingPayments)
	if result.Error != nil {
		return []responses.BookingPaymentResponse{}
	}

	var response []response.BookingPaymentResponse
	for _, bookingPayment := range bookingPayments {
		response = append(response, bookingPayment.MapBookingPaymentToResponse())
	}
	return response
}

func (t *BookingPaymentRepositoryImplementation) FindTotalAmountByBookingId(id string) float64 {
	var totalAmount float64
	var bookingPayments []models.BookingPayment

	result := t.Db.Where("booking_id = ?", id).Find(&bookingPayments)
	if result.Error != nil {
		return 0
	}

	for _, bookingPayment := range bookingPayments {
		totalAmount += bookingPayment.PaymentAmount
	}

	return totalAmount
}
