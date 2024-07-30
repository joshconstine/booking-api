package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type EntityBookingDurationRuleServiceImplementation struct {
	EntityBookingDurationRuleRepository repositories.EntityBookingDurationRuleRepository
}

func NewEntityBookingDurationRuleServiceImplementation(entityBookingDurationRuleRepository repositories.EntityBookingDurationRuleRepository) EntityBookingDurationRuleService {
	return &EntityBookingDurationRuleServiceImplementation{
		EntityBookingDurationRuleRepository: entityBookingDurationRuleRepository,
	}
}

func (t *EntityBookingDurationRuleServiceImplementation) FindByID(entityID uint, entityType string) response.EntityBookingDurationRuleResponse {
	return t.EntityBookingDurationRuleRepository.FindById(entityID, entityType)
}

func (t *EntityBookingDurationRuleServiceImplementation) Update(updatedRule request.UpdateEntityBookingDurationRuleRequest) (response.EntityBookingDurationRuleResponse, error) {
	return t.EntityBookingDurationRuleRepository.Update(updatedRule)
}
