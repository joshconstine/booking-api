package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (p *PaymentMethod) TableName() string {
	return "payment_methods"
}

func (p *PaymentMethod) MapPaymentMethodToResponse() response.PaymentMethodResponse {

	response := response.PaymentMethodResponse{
		ID:   p.ID,
		Name: p.Name,
	}

	return response

}
