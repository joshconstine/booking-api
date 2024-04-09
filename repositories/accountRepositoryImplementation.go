package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type AccountRepositoryImplementation struct {
	Db *gorm.DB
}

func NewAccountRepositoryImplementation(db *gorm.DB) AccountRepository {
	return &AccountRepositoryImplementation{
		Db: db,
	}
}

func (repository *AccountRepositoryImplementation) Create(account models.Account) error {
	result := repository.Db.Model(&models.Account{}).Create(&account)

	return result.Error
}

func (repository *AccountRepositoryImplementation) GetByID(id uint) (models.Account, error) {
	var account models.Account
	result := repository.Db.Preload("Members").Preload("Settings").First(&account, id)
	return account, result.Error
}
