package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type AccountRepository interface {
	Create(account models.Account) error
	Update(account models.Account) error
	GetByID(id uint) (response.AccountResponse, error)
}
