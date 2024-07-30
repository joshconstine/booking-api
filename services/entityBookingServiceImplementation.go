package services

import (
	"booking-api/constants"
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
func isEntityBookingInProgress(entityBooking *response.EntityBookingResponse) bool {
	now := time.Now()
	if entityBooking.Timeblock.StartTime.Before(now) && entityBooking.Timeblock.EndTime.After(now) {
		return true
	}
	return false
}
func (e *EntityBookingServiceImplementation) AuditEntityBookingStatusForBooking(bookingInformation *response.BookingInformationResponse, entityBooking *response.EntityBookingResponse) error {

	//check if user has permission to book this entity.
	//return error if not allowed.
	//check if entity can be booked by this user on this day.

	var entityBookingToAudit response.EntityBookingResponse

	for _, e := range bookingInformation.Entities {
		if e.ID == entityBooking.ID {
			entityBookingToAudit = e
		}
	}

	var request request.UpdateEntityBookingStatusRequest
	request.EntityBookingID = entityBookingToAudit.ID

	if bookingInformation.Details.PaymentComplete == true && bookingInformation.Details.DocumentsSigned == true {
		//Check if the booking is in Progress
		//update booking status to in progress

		if isEntityBookingInProgress(entityBooking) {
			entityBooking.Status.ID = constants.BOOKING_STATUS_IN_PROGRESS_ID

			if entityBookingToAudit.Status.ID != constants.BOOKING_STATUS_IN_PROGRESS_ID {
				request.BookingStatusID = constants.BOOKING_STATUS_IN_PROGRESS_ID
				e.entityBookingRepository.UpdateStatus(request)
			}
		} else {

			entityBooking.Status.ID = constants.BOOKING_STATUS_CONFIRMED_ID

			if entityBookingToAudit.Status.ID != constants.BOOKING_STATUS_CONFIRMED_ID {
				request.BookingStatusID = constants.BOOKING_STATUS_CONFIRMED_ID
				e.entityBookingRepository.UpdateStatus(request)
			}
		}

		//update booking status to confirmed

	} else if bookingInformation.Details.PaymentComplete == true && bookingInformation.Details.DocumentsSigned == true && bookingIsInPast(*bookingInformation) {

		entityBooking.Status.ID = constants.BOOKING_STATUS_COMPLETED_ID

		if entityBookingToAudit.Status.ID != constants.BOOKING_STATUS_COMPLETED_ID {
			request.BookingStatusID = constants.BOOKING_STATUS_COMPLETED_ID
			e.entityBookingRepository.UpdateStatus(request)
		}
	}

	return nil
}
