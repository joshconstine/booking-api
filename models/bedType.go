package models

import (
	"gorm.io/gorm"
)

type BedType struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (b *BedType) TableName() string {
	return "bed_types"
}
