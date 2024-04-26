package services

import (
	"booking-api/data/response"
	"booking-api/repositories"
)

type AccountServiceImplementation struct {
	AccountRepository repositories.AccountRepository
}

func NewAccountServiceImplementation(accountRepository repositories.AccountRepository) AccountService {
	return &AccountServiceImplementation{
		AccountRepository: accountRepository,
	}
}

func (t AccountServiceImplementation) GetInquiriesSnapshot(accountID uint) (response.AccountInquiriesSnapshot, error) {
	return t.AccountRepository.GetInquiriesSnapshot(accountID)
}

func (t AccountServiceImplementation) GetMessagesSnapshot(accountID uint) (response.AccountMessagesSnapshot, error) {
	return t.AccountRepository.GetMessagesSnapshot(accountID)
}

func (t AccountServiceImplementation) GetUserAccountRoles(userID string) ([]response.MembershipResponse, error) {
	return t.AccountRepository.GetUserAccountRoles(userID)
}
