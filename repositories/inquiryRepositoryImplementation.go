package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type InquiryRepositoryImplementation struct {
	Db *gorm.DB
}

func NewInquiryRepositoryImplementation(db *gorm.DB) InquiryRepository {
	return &InquiryRepositoryImplementation{
		Db: db,
	}
}

func (iri *InquiryRepositoryImplementation) Create(inquiry models.Inquiry) (models.Inquiry, error) {
	result := iri.Db.Create(&inquiry)
	if result.Error != nil {
		return models.Inquiry{}, result.Error
	}
	return inquiry, nil
}

func (iri *InquiryRepositoryImplementation) GetByID(inquiryID uint) (response.InquiryResponse, error) {
	var inquiry models.Inquiry
	result := iri.Db.
		Preload("BookingRequests.InquiryStatus").
		Preload("BookingRequests").
		Preload("User").
		First(&inquiry, inquiryID)
	if result.Error != nil {
		return response.InquiryResponse{}, result.Error
	}
	return inquiry.MapInquiryToResponse(), nil
}
