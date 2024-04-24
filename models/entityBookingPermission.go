package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type EntityBookingPermission struct {
	gorm.Model
	UserID          string    `gorm:"not null"`
	AccountID       uint      `gorm:"not null"`
	EntityID        uint      `gorm:"not null"`
	EntityType      string    `gorm:"not null"`
	InquiryStatusID uint      `gorm:"not null; default:1"`
	StartTime       time.Time `gorm:"not null"`
	EndTime         time.Time `gorm:"not null"`
	InquiryStatus   InquiryStatus
}

func (entityBookingPermission *EntityBookingPermission) TableName() string {
	return "entity_booking_permissions"
}

func (entityBookingPermission *EntityBookingPermission) MapEntityBookingPermissionToResponse() response.EntityBookingPermissionResponse {
	return response.EntityBookingPermissionResponse{
		ID:        entityBookingPermission.ID,
		UserID:    entityBookingPermission.UserID,
		AccountID: entityBookingPermission.AccountID,
		Entity: response.EntityInfoResponse{
			EntityID:   entityBookingPermission.EntityID,
			EntityType: entityBookingPermission.EntityType,
		},
		StartTime:     entityBookingPermission.StartTime,
		EndTime:       entityBookingPermission.EndTime,
		InquiryStatus: entityBookingPermission.InquiryStatus.MapInquiryStatusToResponse(),
	}
}
