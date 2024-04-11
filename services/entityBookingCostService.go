package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingCostService interface {
	FindAllForEntity(entityID uint, entityType string) []response.EntityBookingCostResponse
	Create(request request.CreateEntityBookingCostRequest) (response.EntityBookingCostResponse, error)
	Update(request request.UpdateEntityBookingCostRequest) (response.EntityBookingCostResponse, error)
	Delete(entityID uint, entityType string, bookingCostTypeID uint) error
}
