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
	BookingStatus    BookingStatus
	Timeblock        EntityTimeblock
	BookingCostItems []BookingCostItem `gorm:"foreignKey:EntityBookingID"`
	Documents        []BookingDocument `gorm:"foreignKey:EntityBookingID"`
}

func (e *EntityBooking) TableName() string {
	return "entity_bookings"
}

func (e *EntityBooking) MapEntityBookingToResponse() response.EntityBookingResponse {

	response := response.EntityBookingResponse{
		ID:         e.ID,
		EntityID:   e.EntityID,
		EntityType: e.EntityType,
		BookingID:  e.BookingID,
		Timeblock:  e.Timeblock.MapTimeblockToResponse(),
		Status: response.BookingStatusResponse{
			ID:   e.BookingStatus.ID,
			Name: e.BookingStatus.Name,
		},
	}

	for _, costItem := range e.BookingCostItems {
		response.CostItems = append(response.CostItems, costItem.MapBookingCostItemToResponse())
	}

	for _, document := range e.Documents {
		response.Documents = append(response.Documents, document.MapBookingDocumentToResponse())
	}

	return response
}

func (e *EntityBooking) MapEntityBookingToEntityInfoResponse() response.EntityInfoResponse {
	return response.EntityInfoResponse{
		EntityID:   e.EntityID,
		EntityType: e.EntityType,
	}
}
