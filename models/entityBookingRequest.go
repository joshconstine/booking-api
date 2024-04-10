package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type EntityBookingRequest struct {
	gorm.Model
	InquiryID       uint   `gorm:"not null"`
	InquiryStatusID uint   `gorm:"not null; default:1"`
	EntityID        uint   `gorm:"primaryKey"`
	EntityType      string `gorm:"primaryKey"`
	StartTime       time.Time
	EndTime         time.Time
	InquiryStatus   InquiryStatus
}

func (entityInquiry *EntityBookingRequest) MapEntityBookingRequestToResponse() response.EntityBookingRequestResponse {
	return response.EntityBookingRequestResponse{
		ID:            entityInquiry.ID,
		InquiryID:     entityInquiry.InquiryID,
		EntityID:      entityInquiry.EntityID,
		EntityType:    entityInquiry.EntityType,
		StartTime:     entityInquiry.StartTime,
		EndTime:       entityInquiry.EndTime,
		InquiryStatus: entityInquiry.InquiryStatus.MapInquiryStatusToResponse(),
	}
}
