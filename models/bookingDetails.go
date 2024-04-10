package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type BookingDetails struct {
	gorm.Model
	BookingID        string `gorm:"unique; not null"`
	PaymentComplete  bool
	DepositPaid      bool
	PaymentDueDate   time.Time
	DocumentsSigned  bool
	BookingStartDate time.Time
	GuestCount       int
	LocationID       uint   `gorm:"not null"`
	InvoiceID        string `gorm:"type:varchar(255);unique"`
}

func (b *BookingDetails) TableName() string {
	return "booking_details"
}

func (b *BookingDetails) MapBookingDetailsToResponse() response.BookingDetailsResponse {

	response := response.BookingDetailsResponse{
		ID:               b.ID,
		PaymentComplete:  b.PaymentComplete,
		DepositPaid:      b.DepositPaid,
		PaymentDueDate:   b.PaymentDueDate,
		DocumentsSigned:  b.DocumentsSigned,
		BookingStartDate: b.BookingStartDate,
		LocationID:       b.LocationID,
		InvoiceID:        b.InvoiceID,
	}

	return response

}
