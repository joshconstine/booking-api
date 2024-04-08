package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type Timeblock struct {
	gorm.Model
	StartTime  time.Time
	EndTime    time.Time
	EntityID   uint   `gorm:"primaryKey"`
	EntityType string `gorm:"primaryKey"`
	BookingID  string
}

func (t *Timeblock) TableName() string {
	return "timeblocks"
}

func (t *Timeblock) MapTimeblockToResponse() response.TimeblockResponse {
	return response.TimeblockResponse{
		ID:        t.ID,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
		BookingID: t.BookingID,
	}
}
