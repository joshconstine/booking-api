package models

import (
	"gorm.io/gorm"
)

type BookingCostType struct {
	gorm.Model
	Name string `json:"name"`
}
