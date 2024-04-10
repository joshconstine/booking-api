package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityBookingRule struct {
	gorm.Model
	EntityID                uint   `gorm:"index:idx_entity,unique"`
	EntityType              string `gorm:"index:idx_entity,unique"`
	AdvertiseAtAllLocations bool   `gorm:"default: false"`
	AllowPets               bool   `gorm:"default: false"`
	AllowInstantBooking     bool   `gorm:"default: false"`
	OfferEarlyCheckIn       bool
}

func (e *EntityBookingRule) TableName() string {
	return "entity_booking_rules"
}

func (e *EntityBookingRule) MapEntityBookingRuleToResponse() response.EntityBookingRuleResponse {

	result := response.EntityBookingRuleResponse{
		ID:                      e.ID,
		AdvertiseAtAllLocations: e.AdvertiseAtAllLocations,
		AllowPets:               e.AllowPets,
		AllowInstantBooking:     e.AllowInstantBooking,
		OfferEarlyCheckIn:       e.OfferEarlyCheckIn,
	}

	return result

}
