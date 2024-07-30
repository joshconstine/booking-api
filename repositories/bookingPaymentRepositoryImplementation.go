package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	responses "booking-api/data/response"
	"booking-api/models"
	"gorm.io/gorm"
)

type BookingPaymentRepositoryImplementation struct {
	BookingCostItemRepository BookingCostItemRepository
	Db                        *gorm.DB
}

func NewBookingPaymentRepositoryImplementation(bookingCostItemRepository BookingCostItemRepository, db *gorm.DB) *BookingPaymentRepositoryImplementation {

	return &BookingPaymentRepositoryImplementation{BookingCostItemRepository: bookingCostItemRepository, Db: db}
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

func (t *BookingPaymentRepositoryImplementation) Create(bookingPayment requests.CreateBookingPaymentRequest) (response.BookingPaymentResponse, error) {
	bookingPaymentModel := models.BookingPayment{
		BookingID:       bookingPayment.BookingID,
		PaymentMethodID: bookingPayment.PaymentMethodID,
		PaymentAmount:   bookingPayment.PaymentAmount,
	}

	result := t.Db.Create(&bookingPaymentModel)
	if result.Error != nil {
		return response.BookingPaymentResponse{}, result.Error
	}

	return bookingPaymentModel.MapBookingPaymentToResponse(), nil
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

func (t *BookingPaymentRepositoryImplementation) FindTotalPaidByBookingId(id string) float64 {
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

func (t *BookingPaymentRepositoryImplementation) FindTotalOutstandingAmountByBookingId(id string) float64 {
	totalPaid := t.FindTotalPaidByBookingId(id)

	return t.BookingCostItemRepository.GetTotalCostItemsForBooking(id) - totalPaid
}

func (t *BookingPaymentRepositoryImplementation) CheckIfPaymentIsCompleted(id string) bool {
	var paymentCompleted bool
	result := t.Db.Model(&models.BookingDetails{}).Where("booking_id = ?", id).Pluck("payment_complete", &paymentCompleted)

	if result.Error != nil {
		return false
	}

	return paymentCompleted
}
