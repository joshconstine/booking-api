package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type InquiryStatus struct {
	gorm.Model
	Name string `gorm:"unique; not null"`
}

func (inquiryStatus *InquiryStatus) TableName() string {
	return "inquiry_statuses"
}

func (inquiryStatus *InquiryStatus) MapInquiryStatusToResponse() response.InquiryStatusResponse {
	return response.InquiryStatusResponse{
		ID:   inquiryStatus.ID,
		Name: inquiryStatus.Name,
	}
}
