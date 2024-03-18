package models

import (
	"gorm.io/gorm"
)

type Boat struct {
	gorm.Model
	Name       string
	Occupancy  int
	MaxWeight  int
	Photos     []BoatPhoto
	Timeblocks []*Timeblock `gorm:"polymorphic:Entity"`
}

func (b *Boat) TableName() string {
	return "boats"
}
