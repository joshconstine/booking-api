package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type EntityTimeblock struct {
	gorm.Model
	StartTime       time.Time `gorm:"not null"`
	EndTime         time.Time `gorm:"not null"`
	EntityID        uint      `gorm:"primaryKey"`
	EntityType      string    `gorm:"primaryKey"`
	EntityBookingID uint
}

func (t *EntityTimeblock) TableName() string {
	return "entity_timeblocks"
}

func (t *EntityTimeblock) MapTimeblockToResponse() response.EntityTimeblockResponse {
	return response.EntityTimeblockResponse{
		ID:              t.ID,
		StartTime:       t.StartTime,
		EndTime:         t.EndTime,
		EntityBookingID: t.EntityBookingID,
	}
}
