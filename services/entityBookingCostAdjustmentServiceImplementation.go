package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type EntityBookingCostAdjustmentServiceImplementation struct {
	EntityBookingCostAdjustmentRepository repositories.EntityBookingCostAdjustmentRepository
}

func NewEntityBookingCostAdjustmentServiceImplementation(entityBookingCostAdjustmentRepository repositories.EntityBookingCostAdjustmentRepository) EntityBookingCostAdjustmentService {
	return &EntityBookingCostAdjustmentServiceImplementation{EntityBookingCostAdjustmentRepository: entityBookingCostAdjustmentRepository}
}

func (e *EntityBookingCostAdjustmentServiceImplementation) FindAllForEntity(entityID uint, entityType string) []response.EntityBookingCostAdjustmentResponse {
	return e.EntityBookingCostAdjustmentRepository.FindAllForEntity(entityID, entityType)
}

func (e *EntityBookingCostAdjustmentServiceImplementation) FindAllForEntityAndRange(entityID uint, entityType string, startDate string, endDate string) []response.EntityBookingCostAdjustmentResponse {
	return e.EntityBookingCostAdjustmentRepository.FindAllForEntityAndRange(entityID, entityType, startDate, endDate)
}

func (e *EntityBookingCostAdjustmentServiceImplementation) Create(request request.CreateEntityBookingCostAdjustmentRequest) (response.EntityBookingCostAdjustmentResponse, error) {
	return e.EntityBookingCostAdjustmentRepository.Create(request)
}

func (e *EntityBookingCostAdjustmentServiceImplementation) Update(request request.UpdateEntityBookingCostAdjustmentRequest) (response.EntityBookingCostAdjustmentResponse, error) {
	return e.EntityBookingCostAdjustmentRepository.Update(request)
}

func (e *EntityBookingCostAdjustmentServiceImplementation) Delete(adjustmentID uint) error {
	return e.EntityBookingCostAdjustmentRepository.Delete(adjustmentID)
}
