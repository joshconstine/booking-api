package models

import (
	"gorm.io/gorm"
)

type BookingCostItem struct {
	gorm.Model
	BookingID         uint
	BookingCostTypeID uint
	Amount            float64
	Booking           Booking
	BookingCostType   BookingCostType
}

func (b *BookingCostItem) TableName() string {
	return "booking_cost_items"
}
