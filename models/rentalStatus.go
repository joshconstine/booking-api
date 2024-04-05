package models

import (
	"gorm.io/gorm"
)

type RentalStatus struct {
	gorm.Model
	RentalID uint `gorm:"uniqueIndex"`
	IsClean  bool
}
