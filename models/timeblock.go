package models

import (
	"gorm.io/gorm"
)

type Timeblock struct {
	gorm.Model
	StartTime  string
	EndTime    string
	EntityID   uint   `gorm:"primary_key"`
	EntityType string `gorm:"primary_key"`
	BookingID  *uint
}
