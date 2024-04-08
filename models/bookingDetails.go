package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingDetails struct {
	gorm.Model
	BookingID        string
	PaymentComplete  bool
	DepositPaid      bool
	PaymentDueDate   time.Time
	DocumentsSigned  bool
	BookingStartDate time.Time
	LocationID       uint
	InvoiceID        *string `gorm:"type:varchar(255);unique"`
}

func (b *BookingDetails) TableName() string {
	return "booking_details"
}
