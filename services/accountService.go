package services

import (
	"booking-api/data/response"
)

type AccountService interface {
	GetUserAccountRoles(userID string) ([]response.MembershipResponse, error)
	GetInquiriesSnapshot(accountID uint) (response.AccountInquiriesSnapshot, error)
	GetMessagesSnapshot(accountID uint) (response.AccountMessagesSnapshot, error)
}
