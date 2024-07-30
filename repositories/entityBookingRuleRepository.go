package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingRuleRepository interface {
	FindByID(entityID uint, entityType string) (response.EntityBookingRuleResponse, error)
	Update(entityBookingRule request.UpdateEntityBookingRuleRequest) (response.EntityBookingRuleResponse, error)
}
