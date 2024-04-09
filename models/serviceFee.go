package models

import "gorm.io/gorm"

type ServiceFee struct {
	gorm.Model
	ServicePlanID         uint
	FeePercentage         float64 `gorm:"not null"`
	AppliesToAllCostTypes bool
	BookingCostTypeID     uint
}
