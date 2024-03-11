package models

import (
	"gorm.io/gorm"
)

type Boat struct {
	gorm.Model
	Name      string
	Occupancy int
	MaxWeight int
}

func (b *Boat) TableName() string {
	return "boats"
}
