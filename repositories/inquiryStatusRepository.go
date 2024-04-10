package repositories

import (
	"booking-api/models"
)

type InquiryStatusRepository interface {
	Create(inquiryStatus models.InquiryStatus) (models.InquiryStatus, error)
	GetAll() ([]models.InquiryStatus, error)
}
