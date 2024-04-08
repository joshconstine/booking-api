package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityBooking struct {
	gorm.Model
	EntityID         uint   `gorm:"primaryKey"`
	EntityType       string `gorm:"primaryKey"`
	BookingID        string
	TimeblockID      uint
	BookingStatusID  uint
	BookingCostItems []BookingCostItem `gorm:"foreignKey:EntityBookingID"`
	Timeblock        Timeblock
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

	for _, costItem := range e.BookingCostItems {
		response.BookingCostItems = append(response.BookingCostItems, costItem.MapBookingCostItemToResponse())
	}

	return response
}
