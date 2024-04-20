package request

import (
	"booking-api/models"
	"time"

	"github.com/google/uuid"
)

type CreateBookingRequest struct {
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	Guests         int
	EntityRequests []EntityBookingRequest
}

type EntityBookingRequest struct {
	EntityID   uint
	EntityType string
	StartDate  time.Time
	EndDate    time.Time
}

func (e *EntityBookingRequest) MapEntityBookingRequestToEntityBooking() models.EntityBooking {
	timeblock := CreateEntityTimeblockRequest{
		EntityType: e.EntityType,
		EntityID:   e.EntityID,
		StartDate:  e.StartDate,
		EndDate:    e.EndDate,
	}
	return models.EntityBooking{
		EntityID:   e.EntityID,
		EntityType: e.EntityType,
		Timeblock:  timeblock.MapCreateEntityTimeblockRequestToEntityTimeblock(),
	}
}
func calculatePaymentDueDate(bookingStartDate time.Time) time.Time {
	//the due date will be 90 days before the booking start date if the startdate is < 90 days from now the due date is now
	dueDate := bookingStartDate.AddDate(0, 0, -90)
	if dueDate.Before(time.Now()) {
		return time.Now()
	}
	return dueDate
}

func (cbr *CreateBookingRequest) MapCreateBookingRequestToBooking() models.Booking {
	var earliestStartDate time.Time

	earliestStartDate = cbr.EntityRequests[0].StartDate

	for _, entityRequest := range cbr.EntityRequests {
		if entityRequest.StartDate.Before(earliestStartDate) {
			earliestStartDate = entityRequest.StartDate
		}
	}

	booking := models.Booking{
		//TODO wtf fix this
		UserID: uuid.New().String(),
		Details: models.BookingDetails{
			GuestCount:       cbr.Guests,
			BookingStartDate: earliestStartDate,
			PaymentDueDate:   calculatePaymentDueDate(earliestStartDate),
			PaymentComplete:  false,
			DepositPaid:      false,
			DocumentsSigned:  false,
		},
	}

	for _, entityRequest := range cbr.EntityRequests {
		booking.Entities = append(booking.Entities, entityRequest.MapEntityBookingRequestToEntityBooking())
	}

	return booking
}
