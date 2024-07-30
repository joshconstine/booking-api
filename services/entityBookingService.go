package services

import (
	"time"

	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingService interface {
	FindAllForEntity(entityType string, entityID uint) []response.EntityBookingResponse
	FindById(id uint) response.EntityBookingResponse
	FindAllForBooking(bookingID string) []response.EntityBookingResponse
	AttemptToCreate(entityBooking request.CreateEntityBookingRequest) (response.EntityBookingResponse, error)
	AttemptToUpdate(entityBooking request.UpdateEntityBookingRequest) (response.EntityBookingResponse, error)
	FindAllForEntityForRange(entityType string, entityID uint, startTime *time.Time, endTime *time.Time) []response.EntityBookingResponse
}
