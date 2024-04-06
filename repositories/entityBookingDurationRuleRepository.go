package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingDurationRuleRepository interface {
	FindById(entityID uint, entityType string) response.EntityBookingDurationRuleResponse
	Update(updatedRule request.UpdateEntityBookingDurationRuleRequest) response.EntityBookingDurationRuleResponse
}
