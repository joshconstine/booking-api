package services

import (
	"booking-api/data/response"
)

type AccountService interface {
	GetInquiriesSnapshot(accountID uint) (response.AccountInquiriesSnapshot, error)
}
