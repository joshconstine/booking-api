package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type InquiryStatusRepositoryImplementation struct {
	Db *gorm.DB
}

func NewInquiryStatusRepositoryImplementation(db *gorm.DB) InquiryStatusRepository {
	return &InquiryStatusRepositoryImplementation{
		Db: db,
	}
}

func (isri *InquiryStatusRepositoryImplementation) Create(inquiryStatus models.InquiryStatus) (models.InquiryStatus, error) {
	result := isri.Db.Create(&inquiryStatus)
	if result.Error != nil {
		return models.InquiryStatus{}, result.Error
	}
	return inquiryStatus, nil
}

func (isri *InquiryStatusRepositoryImplementation) GetAll() ([]models.InquiryStatus, error) {
	var inquiryStatuses []models.InquiryStatus
	result := isri.Db.Find(&inquiryStatuses)
	if result.Error != nil {
		return []models.InquiryStatus{}, result.Error
	}
	return inquiryStatuses, nil
}
