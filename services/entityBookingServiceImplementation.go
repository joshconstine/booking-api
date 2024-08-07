package services

import (
	"booking-api/constants"
	"time"

	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type EntityBookingServiceImplementation struct {
	entityBookingRepository  repositories.EntityBookingRepository
	bookingDetailsRepository repositories.BookingDetailsRepository
}

func NewEntityBookingServiceImplementation(entityBookingRepository repositories.EntityBookingRepository, bookingDetailsRepository repositories.BookingDetailsRepository) EntityBookingService {
	return &EntityBookingServiceImplementation{entityBookingRepository: entityBookingRepository, bookingDetailsRepository: bookingDetailsRepository}
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
	res := e.entityBookingRepository.Create(entityBooking)

	if res.ID != 0 {

		e.bookingDetailsRepository.UpdatePaymentCompleteStatus(res.BookingID, false)
	}
	return res, nil

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
func checkIfBookingIsCompleted(request *response.EntityBookingResponse) bool {
	res := true
	if request.Status.ID != constants.BOOKING_STATUS_COMPLETED_ID {
		if request.Status.ID != constants.BOOKING_STATUS_CANCELLED_ID {
			if request.Timeblock.EndTime.Before(time.Now()) {
				res = true
			} else {
				res = false
			}
		} else {
			res = false
		}
	}
	return res
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
	if checkIfBookingIsCompleted(&entityBookingToAudit) {
		entityBooking.Status.ID = constants.BOOKING_STATUS_COMPLETED_ID
		if bookingInformation.Status.ID != constants.BOOKING_STATUS_COMPLETED_ID {
			request.BookingStatusID = constants.BOOKING_STATUS_COMPLETED_ID
			e.entityBookingRepository.UpdateStatus(request)
		}
	}

	return nil
}
