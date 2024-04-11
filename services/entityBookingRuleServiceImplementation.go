package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type EntityBookingRuleServiceImplementation struct {
	EntityBookingRuleRepository repositories.EntityBookingRuleRepository
}

func NewEntityBookingRuleServiceImplementation(entityBookingRuleRepository repositories.EntityBookingRuleRepository) EntityBookingRuleService {
	return &EntityBookingRuleServiceImplementation{EntityBookingRuleRepository: entityBookingRuleRepository}
}

func (e *EntityBookingRuleServiceImplementation) FindByID(entityID uint, entityType string) (response.EntityBookingRuleResponse, error) {
	return e.EntityBookingRuleRepository.FindByID(entityID, entityType)
}

func (e *EntityBookingRuleServiceImplementation) Update(entityBookingRule request.UpdateEntityBookingRuleRequest) (response.EntityBookingRuleResponse, error) {
	return e.EntityBookingRuleRepository.Update(entityBookingRule)
}
