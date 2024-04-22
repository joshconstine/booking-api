package models

import (
	"time"

	"gorm.io/gorm"
)

type EntityBookingPermission struct {
	gorm.Model
	UserID          uint      `gorm:"not null"`
	AccountID       uint      `gorm:"not null"`
	EntityID        uint      `gorm:"not null"`
	EntityType      string    `gorm:"not null"`
	InquiryStatusID uint      `gorm:"not null; default:1"`
	StartTime       time.Time `gorm:"not null"`
	EndTime         time.Time `gorm:"not null"`
	InquiryStatus   InquiryStatus
}
