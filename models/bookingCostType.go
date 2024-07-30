package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BookingCostType struct {
	gorm.Model
	Name string `gorm:"not null; unique"`
}

func (b *BookingCostType) TableName() string {
	return "booking_cost_types"
}

func (b *BookingCostType) MapBookingCostTypeToResponse() response.BookingCostTypeResponse {

	costTypeResponse := response.BookingCostTypeResponse{
		ID:   b.ID,
		Name: b.Name,
	}

	return costTypeResponse
}
