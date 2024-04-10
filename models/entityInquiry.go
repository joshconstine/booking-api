package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type EntityInquiry struct {
	gorm.Model
	InquiryID       uint   `gorm:"not null"`
	InquiryStatusID uint   `gorm:"not null; default:1"`
	EntityID        uint   `gorm:"primaryKey"`
	EntityType      string `gorm:"primaryKey"`
	StartTime       time.Time
	EndTime         time.Time
	InquiryStatus   InquiryStatus
}

func (entityInquiry *EntityInquiry) MapEntityInquiryToResponse() response.EntityInquiryResponse {
	return response.EntityInquiryResponse{
		ID:            entityInquiry.ID,
		InquiryID:     entityInquiry.InquiryID,
		EntityID:      entityInquiry.EntityID,
		EntityType:    entityInquiry.EntityType,
		StartTime:     entityInquiry.StartTime,
		EndTime:       entityInquiry.EndTime,
		InquiryStatus: entityInquiry.InquiryStatus.MapInquiryStatusToResponse(),
	}
}
