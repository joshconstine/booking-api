package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type EntityBookingCostServiceImplementation struct {
	EntityBookingCostRepository repositories.EntityBookingCostRepository
}

func NewEntityBookingCostServiceImplementation(entityBookingCostRepository repositories.EntityBookingCostRepository) EntityBookingCostService {
	return &EntityBookingCostServiceImplementation{EntityBookingCostRepository: entityBookingCostRepository}
}

func (e *EntityBookingCostServiceImplementation) FindAllForEntity(entityID uint, entityType string) []response.EntityBookingCostResponse {
	return e.EntityBookingCostRepository.FindAllForEntity(entityID, entityType)
}

func (e *EntityBookingCostServiceImplementation) Create(request request.CreateEntityBookingCostRequest) (response.EntityBookingCostResponse, error) {
	return e.EntityBookingCostRepository.Create(request)
}

func (e *EntityBookingCostServiceImplementation) Update(request request.UpdateEntityBookingCostRequest) (response.EntityBookingCostResponse, error) {
	return e.EntityBookingCostRepository.Update(request)
}

func (e *EntityBookingCostServiceImplementation) Delete(entityID uint, entityType string, bookingCostTypeID uint) error {
	return e.EntityBookingCostRepository.Delete(entityID, entityType, bookingCostTypeID)

}
