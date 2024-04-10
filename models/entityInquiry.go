package models

import (
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
}
