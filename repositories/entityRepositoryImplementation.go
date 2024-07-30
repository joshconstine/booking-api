package repositories

import (
	"booking-api/constants"
	"booking-api/data/response"
	"booking-api/models"
	"gorm.io/gorm"
)

type EntityRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityRepositoryImplementation(db *gorm.DB) EntityRepository {
	return &EntityRepositoryImplementation{Db: db}
}

func (e *EntityRepositoryImplementation) IsUserAdminOfEntity(userID string, memberships []response.MembershipResponse, entityType string, entityID uint) (bool, error) {

	var accountID uint

	switch entityType {
	case constants.BOAT_ENTITY:
		boat := models.Boat{}
		query := e.Db.Model(&models.Boat{}).Where("id = ?", entityID).Select("account_id").First(&boat)
		if query.Error != nil {
			return false, query.Error
		}
		accountID = boat.AccountID

	case constants.RENTAL_ENTITY:
		rental := models.Rental{}
		query := e.Db.Model(&models.Rental{}).Where("id = ?", entityID).Select("account_id").First(&rental)
		if query.Error != nil {
			return false, query.Error

		}
		accountID = rental.AccountID

	default:
		return false, nil
	}

	for _, membership := range memberships {
		if membership.AccountID == accountID && (membership.Role.ID == constants.USER_ROLE_ACCOUNT_OWNER_ID || membership.Role.ID == constants.USER_ROLE_ADMIN_ID) {
			return true, nil
		}

	}

	return false, nil
}
