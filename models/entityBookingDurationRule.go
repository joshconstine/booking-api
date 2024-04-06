package models

import (
	"time"

	"gorm.io/gorm"
)

type EntityBookingDurationRule struct {
	gorm.Model
	EntityID        uint   `gorm:"index:idx_entity,unique"`
	EntityType      string `gorm:"index:idx_entity,unique"`
	MinimumDuration int
	MaximumDuration int
	BookingBuffer   int
	StartTime       time.Time // the time of day that the booking can start
	EndTime         time.Time // the time of day that the booking will end
}
