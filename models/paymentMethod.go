package models

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name string
}

func (p *PaymentMethod) TableName() string {
	return "payment_methods"
}
