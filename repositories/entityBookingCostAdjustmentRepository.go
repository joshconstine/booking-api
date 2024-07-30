package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingCostAdjustmentRepository interface {
	FindAllForEntity(entityID uint, entityType string) []response.EntityBookingCostAdjustmentResponse
	FindAllForEntityAndRange(entityID uint, entityType string, startDate string, endDate string) []response.EntityBookingCostAdjustmentResponse
	Create(request request.CreateEntityBookingCostAdjustmentRequest) (response.EntityBookingCostAdjustmentResponse, error)
	Update(request request.UpdateEntityBookingCostAdjustmentRequest) (response.EntityBookingCostAdjustmentResponse, error)
	Delete(adjustmentID uint) error
}
