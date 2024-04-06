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

	result := t.Db.Find(&bookingPayments)
	if result.Error != nil {
		return []responses.BookingPaymentResponse{}
	}

	var item responses.BookingPaymentResponse
	for _, bookingPayment := range bookingPayments {
		item.ID = bookingPayment.ID
		item.BookingID = bookingPayment.BookingID
		item.PaymentMethod.ID = bookingPayment.PaymentMethodID
		item.PaymentMethod.Name = bookingPayment.PaymentMethod.Name
		item.PaymentAmount = bookingPayment.PaymentAmount

		response = append(response, item)

	}
	return response
}

func (t *BookingPaymentRepositoryImplementation) FindById(id uint) response.BookingPaymentResponse {
	var bookingPayment models.BookingPayment
	var response response.BookingPaymentResponse

	result := t.Db.First(&bookingPayment, id)
	if result.Error != nil {

		return response
	}

	response.ID = bookingPayment.ID
	response.BookingID = bookingPayment.BookingID
	response.PaymentMethod.ID = bookingPayment.PaymentMethodID
	response.PaymentMethod.Name = bookingPayment.PaymentMethod.Name
	response.PaymentAmount = bookingPayment.PaymentAmount

	return response
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

	return response.BookingPaymentResponse{
		ID:            bookingPaymentModel.ID,
		BookingID:     bookingPaymentModel.BookingID,
		PaymentMethod: response.PaymentMethodResponse{ID: bookingPaymentModel.PaymentMethodID, Name: bookingPaymentModel.PaymentMethod.Name},
		PaymentAmount: bookingPaymentModel.PaymentAmount,
		PaymentDate:   bookingPaymentModel.Model.CreatedAt,
	}
}
func (t *BookingPaymentRepositoryImplementation) FindByBookingId(id string) []response.BookingPaymentResponse {
	var bookingPayments []models.BookingPayment
	var response []response.BookingPaymentResponse

	result := t.Db.Where("booking_id = ?", id).Find(&bookingPayments)
	if result.Error != nil {
		return []responses.BookingPaymentResponse{}
	}

	var item responses.BookingPaymentResponse
	for _, bookingPayment := range bookingPayments {
		item.ID = bookingPayment.ID
		item.BookingID = bookingPayment.BookingID
		item.PaymentMethod.ID = bookingPayment.PaymentMethodID
		item.PaymentMethod.Name = bookingPayment.PaymentMethod.Name
		item.PaymentAmount = bookingPayment.PaymentAmount

		response = append(response, item)

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
