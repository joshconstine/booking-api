package request

import (
	"booking-api/models"
	"log"
	"time"
)

type CreateBookingRequest struct {
	InquiryID      uint
	UserID         string
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	Guests         int
	EntityRequests []BookEntityRequest
}

type CreateBookingWithUserInformationRequest struct {
	FirstName string
	LastName  string
	Email     string
}

type BookEntityRequest struct {
	EntityID   uint
	EntityType string
	StartTime  time.Time
	EndTime    time.Time
}

func (e *BookEntityRequest) MapEntityBookingRequestToEntityBooking() models.EntityBooking {
	timeblock := CreateEntityTimeblockRequest{
		EntityType: e.EntityType,
		EntityID:   e.EntityID,
		StartDate:  e.StartTime,
		EndDate:    e.EndTime,
	}

	return models.EntityBooking{
		EntityID:         e.EntityID,
		EntityType:       e.EntityType,
		Timeblock:        timeblock.MapCreateEntityTimeblockRequestToEntityTimeblock(),
		BookingCostItems: []models.BookingCostItem{},
	}

}

func (uir *CreateBookingWithUserInformationRequest) MapCreateBookingWithUserInformationRequestToCreateBookingRequest() CreateBookingRequest {
	return CreateBookingRequest{
		FirstName: uir.FirstName,
		LastName:  uir.LastName,
		Email:     uir.Email,
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
	if cbr.EntityRequests == nil || len(cbr.EntityRequests) == 0 {
		log.Println("EntityRequests is empty")
		return models.Booking{
			UserID:  cbr.UserID,
			Details: models.BookingDetails{},
		}
	}

	log.Printf("Mapping %d entity requests\n", len(cbr.EntityRequests))
	//earliestStartDate := cbr.EntityRequests[0].StartTime

	ebr := cbr.EntityRequests

	earliestStartDate := ebr[0].StartTime
	for _, entityRequest := range cbr.EntityRequests {
		if entityRequest.StartTime.Before(earliestStartDate) {
			earliestStartDate = entityRequest.StartTime
		}
	}

	booking := models.Booking{
		UserID: cbr.UserID,
		Details: models.BookingDetails{
			// Initialize booking details here
			BookingStartDate: earliestStartDate,
		},
	}

	for _, entityRequest := range cbr.EntityRequests {
		booking.Entities = append(booking.Entities, entityRequest.MapEntityBookingRequestToEntityBooking())
	}

	return booking
}
