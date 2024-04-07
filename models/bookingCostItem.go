package models

import (
	"gorm.io/gorm"
)

type BookingCostItem struct {
	gorm.Model
	BookingID         string
	BookingCostTypeID uint
	Amount            float64
	Booking           Booking
	BookingCostType   BookingCostType
	EntityBookingID   uint
}

func (b *BookingCostItem) TableName() string {
	return "booking_cost_items"
}
