package services

import (
	"time"

	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type EntityBookingServiceImplementation struct {
	entityBookingRepository repositories.EntityBookingRepository
}

func NewEntityBookingServiceImplementation(entityBookingRepository repositories.EntityBookingRepository) EntityBookingService {
	return &EntityBookingServiceImplementation{entityBookingRepository: entityBookingRepository}
}

func (e *EntityBookingServiceImplementation) FindAllForEntity(entityType string, entityID uint) []response.EntityBookingResponse {
	return e.entityBookingRepository.FindAllForEntity(entityType, entityID)
}

func (e *EntityBookingServiceImplementation) FindById(id uint) response.EntityBookingResponse {
	return e.entityBookingRepository.FindById(id)
}

func (e *EntityBookingServiceImplementation) FindAllForBooking(bookingID string) []response.EntityBookingResponse {
	return e.entityBookingRepository.FindAllForBooking(bookingID)
}

func (e *EntityBookingServiceImplementation) AttemptToCreate(entityBooking request.CreateEntityBookingRequest) (response.EntityBookingResponse, error) {
	return e.entityBookingRepository.Create(entityBooking), nil
}

func (e *EntityBookingServiceImplementation) AttemptToUpdate(entityBooking request.UpdateEntityBookingRequest) (response.EntityBookingResponse, error) {
	return e.entityBookingRepository.Update(entityBooking), nil
}

func (e *EntityBookingServiceImplementation) FindAllForEntityForRange(entityType string, entityID uint, startTime *time.Time, endTime *time.Time) []response.EntityBookingResponse {
	return e.entityBookingRepository.FindAllForEntityForRange(entityType, entityID, startTime, endTime)
}
func (e *EntityBookingServiceImplementation) AuditEntityBookingStatusForBooking(bookingInformation *response.BookingInformationResponse, entityBooking *response.EntityBookingResponse) error {

	//check if entity is available on this day.
	//return error if not available.
	//check if user has permission to book this entity.
	//return error if not allowed.
	//check if entity can be booked by this user on this day.

	return nil
}
