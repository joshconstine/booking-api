package models

import (
	"gorm.io/gorm"
)

type Postal struct {
	gorm.Model
	PostalCode string `gorm:"not null"`
	CountryIso string `gorm:"not null"`
	Country    Country
}

func (c *Postal) TableName() string {
	return "postals"
}
