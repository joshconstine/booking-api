package models

import (
	"gorm.io/gorm"
)

type Locality struct {
	gorm.Model
	RegionID uint   `gorm:"not null"`
	Name     string `gorm:"not null"`
	Region   Region
}

func (c *Locality) TableName() string {
	return "localities"
}
