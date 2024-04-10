package repositories

import (
	"booking-api/data/response"
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

func (repository *AccountRepositoryImplementation) GetByID(id uint) (response.AccountResponse, error) {
	var account models.Account
	result := repository.Db.
		Preload("Members.Role").
		Preload("Members.User").
		Preload("AccountSettings.AccountOwner").
		Preload("AccountSettings.ServicePlan.Fees").
		First(&account, id)
	return account.MapAccountToResponse(), result.Error
}

func (repository *AccountRepositoryImplementation) Update(account models.Account) error {
	result := repository.Db.Model(&models.Account{}).Save(&account)

	return result.Error
}
