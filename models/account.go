package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name            string          `gorm:"not null; unique"`
	AccountSettings AccountSettings `gorm:"not null"`
	Members         []Membership
	Rentals         []Rental
	Boats           []Boat
}

func (a *Account) TableName() string {
	return "accounts"
}

func (a *Account) MapAccountToResponse() response.AccountResponse {
	return response.AccountResponse{
		ID:              a.ID,
		Name:            a.Name,
		Members:         MapMembershipsToResponses(a.Members),
		AccountSettings: a.AccountSettings.MapAccountSettingsToResponse(),
	}
}
