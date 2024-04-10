package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type EntityBookingCostAdjustment struct {
	gorm.Model
	EntityID          uint    `gorm:"index:idx_entity_cost_type; not null"`
	EntityType        string  `gorm:"index:idx_entity_cost_type; not null"`
	BookingCostTypeID uint    `gorm:"index:idx_entity_cost_type: not null"`
	Amount            float64 `gorm:"not null"`
	TaxRateID         uint    `gorm:"not null"`
	BookingCostType   BookingCostType
	TaxRate           TaxRate
	StartDate         time.Time `gorm:"not null"`
	EndDate           time.Time `gorm:"not null"`
}

func (e *EntityBookingCostAdjustment) TableName() string {
	return "entity_booking_cost_adjustments"
}

func (e *EntityBookingCostAdjustment) MapEntityBookingCostAdjustmentToResponse() response.EntityBookingCostAdjustmentResponse {
	result := response.EntityBookingCostAdjustmentResponse{
		ID:        e.ID,
		Amount:    e.Amount,
		StartDate: e.StartDate,
		EndDate:   e.EndDate,
	}

	result.BookingCostType = e.BookingCostType.MapBookingCostTypeToResponse()
	result.TaxRate = e.TaxRate.MapTaxRateToResponse()

	return result

}
