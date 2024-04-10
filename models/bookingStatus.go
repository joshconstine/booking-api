package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BookingStatus struct {
	gorm.Model
	Name string `gorm:"unique; not null"`
}

func (b *BookingStatus) TableName() string {
	return "booking_statuses"
}

func (b *BookingStatus) MapBookingStatusToResponse() response.BookingStatusResponse {

	response := response.BookingStatusResponse{
		ID:   b.ID,
		Name: b.Name,
	}

	return response

}
