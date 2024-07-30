package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BookingCostItem struct {
	gorm.Model
	BookingID         string  `gorm:"not null"`
	BookingCostTypeID uint    `gorm:"not null"`
	Amount            float64 `gorm:"not null"`
	EntityBookingID   uint
	TaxRateID         uint `gorm:"not null"`
	BookingCostType   BookingCostType
	Booking           Booking
	TaxRate           TaxRate
}

func (b *BookingCostItem) TableName() string {
	return "booking_cost_items"
}

func (b *BookingCostItem) MapBookingCostItemToResponse() response.BookingCostItemResponse {

	result := response.BookingCostItemResponse{
		ID:     b.ID,
		Amount: b.Amount,
	}

	result.BookingCostType = b.BookingCostType.MapBookingCostTypeToResponse()
	result.TaxRate = b.TaxRate.MapTaxRateToResponse()

	return result

}
