package repositories

import (
	"booking-api/constants"
	"booking-api/data/response"
	"booking-api/models"
	"fmt"

	"gorm.io/gorm"
)

type AccountRepositoryImplementation struct {
	Db                              *gorm.DB
	bookingRepositoryImplementation BookingRepository
}

func NewAccountRepositoryImplementation(db *gorm.DB, bookingRepositoryImplementation BookingRepository) *AccountRepositoryImplementation {
	return &AccountRepositoryImplementation{Db: db, bookingRepositoryImplementation: bookingRepositoryImplementation}
}

func GetEntityNameFromIDAndType(db *gorm.DB, id uint, entityType string) string {
	var name string
	switch entityType {
	case constants.RENTAL_ENTITY:
		result := db.Model(&models.Rental{}).Where("id = ?", id).Pluck("name", &name)
		if result.Error != nil {
			return ""
		}

	case constants.BOAT_ENTITY:
		result := db.Model(&models.Boat{}).Where("id = ?", id).Pluck("name", &name)
		if result.Error != nil {
			return ""
		}
	default:

	}

	return name

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

func (repository *AccountRepositoryImplementation) GetInquiriesSnapshot(accountID uint) (snap response.AccountInquiriesSnapshot, err error) {

	var boookingPermissionRequests []models.EntityBookingPermission
	var inqSnapResponse response.InquirySnapshotResponse
	var notifications int64
	result := repository.Db.Model(&models.EntityBookingPermission{}).Where("account_id = ? AND inquiry_status_id = ?", accountID, constants.INQUIRY_STATUS_NEW_ID).Order("created_at desc").Preload("InquiryStatus").Limit(5).Find(&boookingPermissionRequests)

	if result.Error != nil {
		return snap, result.Error
	}

	result = repository.Db.Model(&models.EntityBookingPermission{}).Where("account_id = ? AND inquiry_status_id = ?", accountID, constants.INQUIRY_STATUS_NEW_ID).Where("inquiry_status_id = ?", constants.INQUIRY_STATUS_NEW_ID).Count(&notifications)

	if result.Error != nil {
		return snap, result.Error
	}

	snap.Notifications = uint(notifications)

	var ebpr response.EntityBookingPermissionResponse
	for _, permission := range boookingPermissionRequests {

		ebpr = permission.MapEntityBookingPermissionToResponse()
		ebpr.Entity.Name = GetEntityNameFromIDAndType(repository.Db, permission.EntityID, permission.EntityType)
		inqSnapResponse.PermissionRequests = append(inqSnapResponse.PermissionRequests, ebpr)
		inqSnapResponse.Chat = response.ChatSnapshotResponse{}
		snap.Inquiries = append(snap.Inquiries, inqSnapResponse)
		inqSnapResponse = response.InquirySnapshotResponse{}
		ebpr = response.EntityBookingPermissionResponse{}

	}

	return snap, nil
}
func (repository *AccountRepositoryImplementation) GetMessagesSnapshot(accountID uint) (snap response.AccountMessagesSnapshot, err error) {
	var chats []models.Chat
	var notifications int64

	// result := repository.Db.Model(&models.Chat{}).Where("account_id = ?", accountID).Order("created_at desc").Limit(5).Find(&chats)
	result := repository.Db.Model(&models.Chat{}).
		// Joins("JOIN users ON users.id = chats.user_id").
		Where("account_id = ?", accountID).
		Order("chats.created_at DESC").
		Limit(1).
		Preload("Messages").
		Limit(3).
		Find(&chats)

	if result.Error != nil {
		return snap, result.Error
	}

	result = repository.Db.Model(&models.Chat{}).Where("account_id = ?", accountID).Preload("Messages").Count(&notifications)

	if result.Error != nil {
		return snap, result.Error
	}

	snap.Notifications = uint(notifications)

	var chatSnap response.ChatSnapshotResponse

	for _, chat := range chats {
		chatSnap.ChatID = chat.ID
		if len(chat.Messages) > 0 {
			chatSnap.Message = chat.Messages[0].Message
		}
		// chatSnap.Name =
		snap.Chats = append(snap.Chats, chatSnap)
		chatSnap = response.ChatSnapshotResponse{}
	}

	return snap, nil
}

func (repository *AccountRepositoryImplementation) GetUserAccountRoles(userID string) ([]response.MembershipResponse, error) {
	var memberships []models.Membership
	result := repository.Db.
		Preload("Role").
		Where("user_id = ?", userID).
		Find(&memberships)

	return models.MapMembershipsToResponses(memberships), result.Error
}

func (repository *AccountRepositoryImplementation) GetAccountSettings(accountID uint) (response.AccountSettingsResponse, error) {
	var accountSettings models.AccountSettings
	result := repository.Db.
		Preload("ServicePlan.Fees").
		Preload("AccountOwner").
		First(&accountSettings, "account_id = ?", accountID)

	return accountSettings.MapAccountSettingsToResponse(), result.Error
}

func (repository *AccountRepositoryImplementation) AddStripeIDToAccountSettings(accountID uint, stripeAccountID string) error {
	var accountSettings models.AccountSettings
	accountSettings.AccountID = accountID
	accountSettings.StripeAccountID = stripeAccountID

	result := repository.Db.
		FirstOrCreate(&accountSettings, "account_id = ?", accountID)

	if result.Error != nil {
		return result.Error
	}
	// if accountSettings.StripeAccountID != "" {
	// 	return nil
	// }

	// accountSettings.StripeAccountID = stripeAccountID
	// result = repository.Db.Save(&accountSettings)

	return result.Error
}

//	func (repository *AccountRepositoryImplementation) GetAccountIDForBooking(bookingID string) (string, error) {
//		var booking models.Booking
//		result := repository.Db.Model(&models.Booking{}).Preload("Entities").First(&booking, bookingID)
//
//		if result.Error != nil {
//			return "", result.Error
//		}
//
//		var accountID uint
//
//		for _, entity := range booking.Entities {
//			if entity.EntityType == constants.RENTAL_ENTITY {
//				rentalresult := repository.Db.Model(&models.Rental{}).Where("id = ?", entity.EntityID).Pluck("account_id", &accountID)
//				if rentalresult.Error != nil {
//					return "error getting rental details", rentalresult.Error
//				}
//			} else if entity.EntityType == constants.BOAT_ENTITY {
//				boatresult := repository.Db.Model(&models.Boat{}).Where("id = ?", entity.EntityID).Pluck("account_id", &accountID)
//				if boatresult.Error != nil {
//					return "error reading boat details", boatresult.Error
//				}
//
//			}
//		}
//		var stripeAccountID string
//
//		stripeAccountResult := repository.Db.Model(&models.AccountSettings{}).Where("account_id = ?", accountID).Pluck("stripe_account_id", &stripeAccountID)
//
//		if stripeAccountResult.Error != nil {
//			return "error reading stripe account id for account", stripeAccountResult.Error
//		}
//		return stripeAccountID, nil
//
// }
func (repository *AccountRepositoryImplementation) GetAccountIDForBooking(bookingID string) (string, error) {

	var booking response.BookingInformationResponse

	booking = repository.bookingRepositoryImplementation.FindById(bookingID)

	var accountID uint
	var found bool

	for _, entity := range booking.Entities {
		if entity.EntityType == constants.RENTAL_ENTITY {
			rentalResult := repository.Db.Model(&models.Rental{}).Where("id = ?", entity.EntityID).Pluck("account_id", &accountID)
			if rentalResult.Error != nil {
				return "", fmt.Errorf("error getting rental details for entity ID %v: %w", entity.EntityID, rentalResult.Error)
			}
			found = true
		} else if entity.EntityType == constants.BOAT_ENTITY {
			boatResult := repository.Db.Model(&models.Boat{}).Where("id = ?", entity.EntityID).Pluck("account_id", &accountID)
			if boatResult.Error != nil {
				return "", fmt.Errorf("error getting boat details for entity ID %v: %w", entity.EntityID, boatResult.Error)
			}
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("no rental or boat entity found for booking ID %v", bookingID)
	}

	var stripeAccountID string
	stripeAccountResult := repository.Db.Model(&models.AccountSettings{}).Where("account_id = ?", accountID).Pluck("stripe_account_id", &stripeAccountID)
	if stripeAccountResult.Error != nil {
		return "", fmt.Errorf("error reading stripe account ID for account ID %v: %w", accountID, stripeAccountResult.Error)
	}

	return stripeAccountID, nil
}
