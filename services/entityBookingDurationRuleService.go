package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingDurationRuleService interface {
	FindByID(entityID uint, entityType string) response.EntityBookingDurationRuleResponse
	Update(updatedRule request.UpdateEntityBookingDurationRuleRequest) (response.EntityBookingDurationRuleResponse, error)
}
