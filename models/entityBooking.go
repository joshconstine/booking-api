package models

import (
	"gorm.io/gorm"
)

type EntityBooking struct {
	gorm.Model
	EntityID         uint   `gorm:"primaryKey"`
	EntityType       string `gorm:"primaryKey"`
	BookingID        string
	TimeblockID      uint
	BookingStatusID  uint
	BookingCostItems []BookingCostItem `gorm:"foreignKey:EntityBookingID"`
	Timeblock        Timeblock
}
