package models

import (
	"gorm.io/gorm"
)

type BoatDefaultSettings struct {
	gorm.Model
	BoatID                  uint
	DailyCost               float64
	MinimumBookingDuration  uint
	AdvertiseAtAllLocations bool
	FileID                  *uint
}

func (b *BoatDefaultSettings) TableName() string {
	return "boat_default_settings"
}
