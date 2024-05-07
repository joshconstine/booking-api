package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type AccountRepository interface {
	GetUserAccountRoles(userID string) ([]response.MembershipResponse, error)
	Create(account models.Account) error
	Update(account models.Account) error
	GetByID(id uint) (response.AccountResponse, error)
	GetInquiriesSnapshot(accountID uint) (response.AccountInquiriesSnapshot, error)
	GetMessagesSnapshot(accountID uint) (response.AccountMessagesSnapshot, error)
	GetAccountSettings(accountID uint) (response.AccountSettingsResponse, error)
	AddStripeIDToAccountSettings(accountID uint, stripeAccountID string) error
}
