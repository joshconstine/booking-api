package models

import "gorm.io/gorm"

type AccountSettings struct {
	gorm.Model
	AccountID      uint `gorm:"not null"`
	PlanLevelID    uint `gorm:"not null"`
	AccountOwnerID uint `gorm:"not null"`
	AccountOwner   Membership
	PlanLevel      ServicePlan
}
