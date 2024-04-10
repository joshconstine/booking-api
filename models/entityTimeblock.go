package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type EntityTimeblock struct {
	gorm.Model
	StartTime  time.Time
	EndTime    time.Time
	EntityID   uint   `gorm:"primaryKey"`
	EntityType string `gorm:"primaryKey"`
	BookingID  string
}

func (t *EntityTimeblock) TableName() string {
	return "entity_timeblocks"
}

func (t *EntityTimeblock) MapTimeblockToResponse() response.EntityTimeblockResponse {
	return response.EntityTimeblockResponse{
		ID:        t.ID,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
		BookingID: t.BookingID,
	}
}
