package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingRuleService interface {
	FindByID(entityID uint, entityType string) (response.EntityBookingRuleResponse, error)
	Update(entityBookingRule request.UpdateEntityBookingRuleRequest) (response.EntityBookingRuleResponse, error)
}
