package repositories

import (
	"booking-api/models"
)

type InquiryRepository interface {
	Create(inquiry models.Inquiry) (models.Inquiry, error)
	GetByID(inquiryID uint) (models.Inquiry, error)
}
