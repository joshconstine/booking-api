package models

import (
	"gorm.io/gorm"
)

type BookingPayment struct {
	gorm.Model
	BookingID       string
	PaymentMethodID uint
	PaypalReference *string
	PaymentAmount   float64
	Booking         Booking
	PaymentMethod   PaymentMethod
}

func (b *BookingPayment) TableName() string {
	return "booking_payments"
}
