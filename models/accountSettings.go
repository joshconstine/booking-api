package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type AccountSettings struct {
	gorm.Model
	AccountID      uint   `gorm:"not null; uniqueIndex"`
	ServicePlanID  uint   `gorm:"not null"`
	AccountOwnerID string `gorm:"not null"`
	AccountOwner   Membership
	ServicePlan    ServicePlan
}

func (a *AccountSettings) TableName() string {
	return "account_settings"
}

func (a *AccountSettings) MapAccountSettingsToResponse() response.AccountSettingsResponse {
	return response.AccountSettingsResponse{
		ID:           a.ID,
		AccountID:    a.AccountID,
		ServicePlan:  a.ServicePlan.MapServicePlanToResponse(),
		AccountOwner: a.AccountOwner.MapMembershipToResponse(),
	}
}
