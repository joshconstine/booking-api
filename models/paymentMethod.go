package models

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (p *PaymentMethod) TableName() string {
	return "payment_methods"
}
