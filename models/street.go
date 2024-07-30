package models

import (
	"gorm.io/gorm"
)

type Street struct {
	gorm.Model
	Name       string `gorm:"not null"`
	LocalityID uint   `gorm:"not null"`
	Locality   Locality
}

func (c *Street) TableName() string {
	return "streets"
}
