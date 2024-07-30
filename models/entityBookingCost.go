package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityBookingCost struct {
	gorm.Model
	EntityID          uint    `gorm:"index:idx_entity_cost_type,unique;"`
	EntityType        string  `gorm:"index:idx_entity_cost_type,unique;"`
	BookingCostTypeID uint    `gorm:"index:idx_entity_cost_type,unique;"`
	Amount            float64 `gorm:"not null"`
	TaxRateID         uint    `gorm:"not null"`
	BookingCostType   BookingCostType
	TaxRate           TaxRate
}

func (e *EntityBookingCost) TableName() string {
	return "entity_booking_costs"
}

func (e *EntityBookingCost) MapEntityBookingCostToResponse() response.EntityBookingCostResponse {

	result := response.EntityBookingCostResponse{
		ID:     e.ID,
		Amount: e.Amount,
	}

	result.BookingCostType = e.BookingCostType.MapBookingCostTypeToResponse()
	result.TaxRate = e.TaxRate.MapTaxRateToResponse()

	return result

}

//When an entity is booked ie a Rental. This function maps the entity booking cost to a booking cost item
// from here we can calculate the total cost of the booking

func (e *EntityBookingCost) MapEntityBookingCostToBookingCostItem(bookingID string, entityBookingId uint) BookingCostItem {

	result := BookingCostItem{
		EntityBookingID:   entityBookingId,
		BookingID:         bookingID,
		Amount:            e.Amount,
		TaxRateID:         e.TaxRateID,
		BookingCostTypeID: e.BookingCostTypeID,
		TaxRate: TaxRate{
			Model: gorm.Model{
				ID: e.TaxRate.ID,
			},
		},
		BookingCostType: BookingCostType{
			Model: gorm.Model{
				ID: e.BookingCostType.ID,
			},
		},
	}

	return result
}
