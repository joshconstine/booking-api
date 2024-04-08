package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BookingCostItem struct {
	gorm.Model
	BookingID         string
	BookingCostTypeID uint
	Amount            float64
	Booking           Booking
	BookingCostType   BookingCostType
	EntityBookingID   uint
}

func (b *BookingCostItem) TableName() string {
	return "booking_cost_items"
}

func (b *BookingCostItem) MapBookingCostItemToResponse() response.BookingCostItemResponse {

	costTypeResponse := b.BookingCostType.MapBookingCostTypeToResponse()

	result := response.BookingCostItemResponse{
		ID:              b.ID,
		BookingID:       b.BookingID,
		Amount:          b.Amount,
		BookingCostType: costTypeResponse,
	}

	return result

}
