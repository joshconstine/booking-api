package models

import (
	"gorm.io/gorm"
)

type BookingStatus struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func (b *BookingStatus) TableName() string {
	return "booking_statuses"
}
