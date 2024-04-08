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
	EntityBookingID   uint
	TaxRateID         uint
	BookingCostType   BookingCostType
	TaxRate           TaxRate
}

func (b *BookingCostItem) TableName() string {
	return "booking_cost_items"
}

func (b *BookingCostItem) MapBookingCostItemToResponse() response.BookingCostItemResponse {

	result := response.BookingCostItemResponse{
		ID:        b.ID,
		BookingID: b.BookingID,
		Amount:    b.Amount,
	}

	result.BookingCostType = b.BookingCostType.MapBookingCostTypeToResponse()
	result.TaxRate = b.TaxRate.MapTaxRateToResponse()

	return result

}
