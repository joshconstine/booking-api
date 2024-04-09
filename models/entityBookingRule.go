package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityBookingRule struct {
	gorm.Model
	EntityID                uint
	EntityType              string
	AdvertiseAtAllLocations bool
	AllowPets               bool
	AllowInstantBooking     bool
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
