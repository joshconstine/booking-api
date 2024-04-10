package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BookingPayment struct {
	gorm.Model
	BookingID       string  `gorm:"not null"`
	PaymentAmount   float64 `gorm:"not null"`
	PaymentMethodID uint    `gorm:"not null"`
	PaypalReference *string
	Booking         Booking
	PaymentMethod   PaymentMethod
}

func (b *BookingPayment) TableName() string {
	return "booking_payments"
}

func (b *BookingPayment) MapBookingPaymentToResponse() response.BookingPaymentResponse {

	result := response.BookingPaymentResponse{
		ID:            b.ID,
		BookingID:     b.BookingID,
		PaymentAmount: b.PaymentAmount,
		PaymentDate:   b.CreatedAt,
	}

	result.PaymentMethod = b.PaymentMethod.MapPaymentMethodToResponse()

	return result

}
