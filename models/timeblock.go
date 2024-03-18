package models

import (
	"gorm.io/gorm"
)

type Timeblock struct {
	gorm.Model
	StartTime  string
	EndTime    string
	EntityID   uint
	EntityType string
	BookingID  *uint
}
