package models

import (
	"gorm.io/gorm"
)

type EntityBookingCost struct {
	gorm.Model
	EntityID          uint   `gorm:"index:idx_entity_cost_type,unique"`
	EntityType        string `gorm:"index:idx_entity_cost_type,unique"`
	BookingCostTypeID uint   `gorm:"index:idx_entity_cost_type,unique"`
	Amount            float64
	TaxRateID         uint
	BookingCostType   BookingCostType
	TaxRate           TaxRate
}

func (e *EntityBookingCost) TableName() string {
	return "entity_booking_costs"
}
