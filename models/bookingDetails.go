package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingDetails struct {
	gorm.Model
	BookingID        int
	PaymentComplete  bool
	PaymentDueDate   time.Time
	DocumentsSigned  bool
	BookingStartDate time.Time
	InvoiceID        string `gorm:"type:varchar(255);unique"`
}

func (b *BookingDetails) TableName() string {
	return "booking_details"
}
