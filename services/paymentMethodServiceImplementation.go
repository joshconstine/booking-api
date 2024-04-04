package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type PaymentMethodServiceImplementation struct {
	PaymentMethodRepository repositories.PaymentMethodRepository
	Validate                *validator.Validate
}

func NewPaymentMethodServiceImplementation(paymentMethodRepository repositories.PaymentMethodRepository, validate *validator.Validate) PaymentMethodService {
	return &PaymentMethodServiceImplementation{
		PaymentMethodRepository: paymentMethodRepository,
		Validate:                validate,
	}
}

func (t PaymentMethodServiceImplementation) FindAll() []response.PaymentMethodResponse {
	result := t.PaymentMethodRepository.FindAll()

	return result
}

func (t PaymentMethodServiceImplementation) FindById(paymentMethodId uint) response.PaymentMethodResponse {
	paymentMethodData := t.PaymentMethodRepository.FindById(paymentMethodId)

	return paymentMethodData
}
