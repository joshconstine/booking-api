package repositories

import "booking-api/models"

type AccountRepository interface {
	Create(account models.Account) error
	GetByID(id uint) (models.Account, error)
}
