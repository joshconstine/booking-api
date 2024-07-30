package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"time"
)

type EntityBookingRepository interface {
	FindAllForEntity(entityType string, entityID uint) []response.EntityBookingResponse
	FindById(id uint) response.EntityBookingResponse
	FindAllForBooking(bookingID string) []response.EntityBookingResponse
	Create(entityBooking request.CreateEntityBookingRequest) response.EntityBookingResponse
	Update(entityBooking request.UpdateEntityBookingRequest) response.EntityBookingResponse
	UpdateStatus(request request.UpdateEntityBookingStatusRequest) error
	FindAllForEntityForRange(entityType string, entityID uint, startTime *time.Time, endTime *time.Time) []response.EntityBookingResponse
}
