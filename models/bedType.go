package models

import (
	"gorm.io/gorm"
)

type BedType struct {
	gorm.Model
	Name string
}

func (b *BedType) TableName() string {
	return "bed_types"
}
