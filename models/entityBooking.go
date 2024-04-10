package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityBooking struct {
	gorm.Model
	EntityID         uint   `gorm:"primaryKey"`
	EntityType       string `gorm:"primaryKey"`
	BookingID        string `gorm:"not null"`
	TimeblockID      uint   `gorm:"not null"`
	BookingStatusID  uint   `gorm:"not null:default:1"`
	Timeblock        EntityTimeblock
	BookingCostItems []BookingCostItem `gorm:"foreignKey:EntityBookingID"`
	Documents        []BookingDocument `gorm:"foreignKey:EntityBookingID"`
}

func (e *EntityBooking) TableName() string {
	return "entity_bookings"
}

func (e *EntityBooking) MapEntityBookingToResponse() response.EntityBookingResponse {

	response := response.EntityBookingResponse{
		ID:              e.ID,
		BookingID:       e.BookingID,
		TimeblockID:     e.TimeblockID,
		BookingStatusID: e.BookingStatusID,
	}

	return response
}
