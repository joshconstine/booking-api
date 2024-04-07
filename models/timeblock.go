package models

import (
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
