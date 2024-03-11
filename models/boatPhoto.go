package models

import (
	"gorm.io/gorm"
)

type BoatPhoto struct {
	gorm.Model
	BoatID   int
	PhotoURL string
	Boat     Boat
}

func (bp *BoatPhoto) TableName() string {
	return "boat_photos"
}
