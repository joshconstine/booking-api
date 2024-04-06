package models

import (
	"gorm.io/gorm"
)

type Timeblock struct {
	gorm.Model
	StartTime  string
	EndTime    string
	EntityID   uint   `gorm:"primaryKey"`
	EntityType string `gorm:"primaryKey"`
	BookingID  string
}

func (t *Timeblock) TableName() string {
	return "timeblocks"
}
