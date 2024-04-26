package repositories

import (
	"booking-api/constants"
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
	result := repository.Db.Model(&models.EntityBookingPermission{}).Where("account_id = ?", accountID).Order("created_at desc").Preload("InquiryStatus").Limit(5).Find(&boookingPermissionRequests)

	if result.Error != nil {
		return snap, result.Error
	}

	result = repository.Db.Model(&models.EntityBookingPermission{}).Where("account_id = ?", accountID).Where("inquiry_status_id = ?", constants.INQUIRY_STATUS_NEW_ID).Count(&notifications)

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
