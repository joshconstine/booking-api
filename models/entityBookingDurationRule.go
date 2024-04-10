package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type EntityBookingDurationRule struct {
	gorm.Model
	EntityID        uint      `gorm:"index:idx_entity,unique; mot null"`
	EntityType      string    `gorm:"index:idx_entity,unique; mot null"`
	MinimumDuration int       `gorm:"not null: default: 1"`
	MaximumDuration int       `gorm:"not null: default: 15"`
	BookingBuffer   int       `gorm:"not null: default: 1"`
	StartTime       time.Time `gorm:"not null: default: Time.Now()"`
	EndTime         time.Time `gorm:"not null: default: Time.Now()"`
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
