package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type PaymentMethodRepository interface {
	FindAll() []response.PaymentMethodResponse
	FindById(id uint) response.PaymentMethodResponse
	Create(paymentMethod requests.CreatePaymentMethodRequest) response.PaymentMethodResponse
}
