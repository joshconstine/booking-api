package repositories

import (
	"booking-api/constants"
	"booking-api/data/response"
	"booking-api/models"

	"time"

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

func (repository *AccountRepositoryImplementation) GetInquiriesSnapshot(accountID uint) (snap response.AccountInquiriesSnapshot, err error) {

	snap.Notifications = 5
	snap.Inquiries = []response.InquirySnapshotResponse{
		{
			Chat: response.ChatSnapshotResponse{
				ChatID:  1,
				Message: "Hello, does the Morey have enough space for 10 people?",
				Name:    "John Doe",
				Sent:    time.Now().Format("2006-01-02 15:04:05"),
			},
			PermissionRequests: []response.EntityBookingPermissionResponse{
				{
					ID:        1,
					AccountID: 1,

					UserID: "1",
					Entity: response.EntityInfoResponse{

						EntityID:   1,
						EntityType: "rental",
						Name:       "The Morey",
					},
					InquiryStatus: response.InquiryStatusResponse{
						ID:   constants.INQUIRY_STATUS_NEW_ID,
						Name: constants.INQUIRY_STATUS_NEW_NAME,
					},
					StartTime: time.Now(),
					EndTime:   time.Now(),
				},
			},
		},
	}

	return snap, nil
}
func (repository *AccountRepositoryImplementation) GetMessagesSnapshot(accountID uint) (snap response.AccountMessagesSnapshot, err error) {

	snap.Notifications = 5
	snap.Chats = []response.ChatSnapshotResponse{
		{
			ChatID:  1,
			Message: "Hello, If I have a wedding here, will the bar be open?",
			Name:    "Jane Do",
			Sent:    time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	return snap, nil
}
