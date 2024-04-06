package models

import (
	"gorm.io/gorm"
)

type Boat struct {
	gorm.Model
	Name       string
	Occupancy  int
	MaxWeight  int
	Timeblocks []*Timeblock  `gorm:"polymorphic:Entity"`
	Photos     []EntityPhoto `gorm:"polymorphic:Entity"`
}

func (b *Boat) TableName() string {
	return "boats"
}
