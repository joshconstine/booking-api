package models

import (
	"booking-api/data/response"
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

func (e *EntityBookingDurationRule) TableName() string {
	return "entity_booking_duration_rules"
}

func (e *EntityBookingDurationRule) MapEntityBookingDurationRuleToResponse() response.EntityBookingDurationRuleResponse {
	return response.EntityBookingDurationRuleResponse{
		ID:              e.ID,
		EntityID:        e.EntityID,
		EntityType:      e.EntityType,
		MinimumDuration: e.MinimumDuration,
		MaximumDuration: e.MaximumDuration,
		Buffer:          e.BookingBuffer,
		StartTime:       e.StartTime,
		EndTime:         e.EndTime,
	}
}
