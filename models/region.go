package models

import (
	"gorm.io/gorm"
)

type Region struct {
	gorm.Model
	Name       string `gorm:"not null"`
	CountryIso string `gorm:"not null"`
	Country    Country
}

func (c *Region) TableName() string {
	return "regions"
}
