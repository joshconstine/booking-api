package models

import (
	"gorm.io/gorm"
)

type BoatStatus struct {
	gorm.Model
	BoatID     uint `gorm:"uniqueIndex"`
	IsClean    bool `gorm:"not null; default:true"`
	LocationID uint `gorm:"not null"`

	Location Location
}

func (b *BoatStatus) TableName() string {
	return "boat_statuses"
}
