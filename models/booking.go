package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID           int
	BookingStatusID  int
	BookingDetailsID int
	User             User
	BookingStatus    BookingStatus
	BookingDetails   BookingDetails
}

func (b *Booking) TableName() string {
	return "bookings"
}
