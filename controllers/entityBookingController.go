package controllers

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/services"
	"booking-api/view/ui"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type EntityBookingController struct {
	entityBookingService services.EntityBookingService
	rentalService        services.RentalService
	boatService          services.BoatService
}

func NewEntityBookingController(entityBookingService services.EntityBookingService, rentalService services.RentalService, boatService services.BoatService) *EntityBookingController {
	return &EntityBookingController{entityBookingService: entityBookingService,
		rentalService: rentalService,
		boatService:   boatService}
}

func (e *EntityBookingController) CreateEntityBooking(w http.ResponseWriter, r *http.Request) error {

	entityType := r.FormValue("entityType")

	params := ui.AddEntityToBookingRequestParams{
		BookingID:  r.FormValue("bookingID"),
		StartTime:  r.FormValue("startTime"),
		EndTime:    r.FormValue("endTime"),
		EntityType: entityType,
	}
	if params.EntityType == constants.BOAT_ENTITY {

		entityIdInt, err := strconv.Atoi(r.FormValue("entityIDBoat"))
		if err != nil {
			return err

		}
		params.EntityID = uint(entityIdInt)
	} else if params.EntityType == constants.RENTAL_ENTITY {

		entityIdInt, err := strconv.Atoi(r.FormValue("entityIDRental"))
		if err != nil {
			return err

		}
		params.EntityID = uint(entityIdInt)
	} else {
		return fmt.Errorf("Invalid entity type")
	}

	// Adjust the time parsing format to match the input
	startTimeTime, err := time.Parse("2006-01-02T15:04", params.StartTime)
	if err != nil {
		return err
	}
	endTimeTime, err := time.Parse("2006-01-02T15:04", params.EndTime)
	if err != nil {
		return err
	}

	rentalEntityBooking := request.CreateEntityBookingRequest{
		BookingID:  params.BookingID,
		EntityID:   params.EntityID,
		EntityType: params.EntityType,
		StartTime:  startTimeTime,
		EndTime:    endTimeTime,
	}

	booking, err := e.entityBookingService.AttemptToCreate(rentalEntityBooking)

	if err != nil {
		return err

	}

	log.Printf("Created entity booking: %v", booking.ID)

	http.Redirect(w, r, fmt.Sprintf("/bookings/%s", params.BookingID), http.StatusFound)
	return nil
}

func (e *EntityBookingController) AddEntityToBookingForm(w http.ResponseWriter, r *http.Request) error {
	bookingID := chi.URLParam(r, "bookingID")
	params := ui.AddEntityToBookingRequestParams{
		BookingID: bookingID,
	}

	rentals := e.rentalService.FindAll()
	boats := e.boatService.FindAll()

	return render(r, w, ui.AddEntityToBookingForm(params, rentals, boats))
}
