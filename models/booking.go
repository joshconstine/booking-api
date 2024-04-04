package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID           uint
	BookingStatusID  uint
	BookingDetailsID uint
	User             User
	BookingStatus    BookingStatus
	BookingDetails   BookingDetails
	BookingCostItems []BookingCostItem
}

func (b *Booking) TableName() string {
	return "bookings"
}
