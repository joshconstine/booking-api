package models

import "gorm.io/gorm"

type ServicePlan struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
	Fees []ServiceFee
}
