package models

import (
	"gorm.io/gorm"
)

type AddressType struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (c *AddressType) TableName() string {
	return "address_types"
}
