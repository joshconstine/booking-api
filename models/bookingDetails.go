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
	InvoiceID        string `gorm:"type:varchar(255)"`
}

func (b *BookingDetails) TableName() string {
	return "booking_details"
}

func (b *BookingDetails) MapBookingDetailsToResponse() response.BookingDetailsResponse {

	response := response.BookingDetailsResponse{
		ID:               b.ID,
		BookingID:        b.BookingID,
		PaymentComplete:  b.PaymentComplete,
		DepositPaid:      b.DepositPaid,
		PaymentDueDate:   b.PaymentDueDate,
		DocumentsSigned:  b.DocumentsSigned,
		BookingStartDate: b.BookingStartDate,
		GuestCount:       b.GuestCount,
		LocationID:       b.LocationID,
		InvoiceID:        b.InvoiceID,
	}

	return response

}
