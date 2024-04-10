package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type InquiryRepository interface {
	Create(inquiry models.Inquiry) (models.Inquiry, error)
	GetByID(inquiryID uint) (response.InquiryResponse, error)
}
