package models

import (
	"gorm.io/gorm"
)

type RentalStatus struct {
	gorm.Model
	RentalID uint
	IsClean  bool
}
