package services

import (
	responses "booking-api/data/response"
)

type PaymentMethodService interface {
	FindAll() []responses.PaymentMethodResponse
	FindById(id uint) responses.PaymentMethodResponse
}
