package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	AddressType   AddressType
	AddressTypeID uint `gorm:"not null"`
	StreetID      uint `gorm:"not null"`
	PostalID      uint `gorm:"not null"`
	LocalityID    uint `gorm:"not null"`
	RegionID      uint `gorm:"not null"`

	Street   Street
	Postal   Postal
	Locality Locality
	Region   Region
	Premise  string
}

func (c *Address) TableName() string {
	return "addresses"

}
