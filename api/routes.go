package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes for the API.
func InitRoutes(r *mux.Router, db *sql.DB) {
	// Define the routes.

	//Bookings
	r.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		GetBookings(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		CreateBooking(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/bookings/{id}/details", func(w http.ResponseWriter, r *http.Request) {
		GetBookingDetails(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings/{id}/rentalBookings", func(w http.ResponseWriter, r *http.Request) {
		GetRentalBookingsForBooking(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings/{id}/boatBookings", func(w http.ResponseWriter, r *http.Request) {
		GetBoatBookingsForBooking(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings/{id}/rentalBookings/{rentalBookingId}/details", func(w http.ResponseWriter, r *http.Request) {
		GetRentalBookingDetails(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings/{id}/boatBookings/{boatBookingId}/details", func(w http.ResponseWriter, r *http.Request) {
		GetBoatBookingDetails(w, r, db)
	}).Methods("GET")

	//Rentals
	r.HandleFunc("/rentals", func(w http.ResponseWriter, r *http.Request) {
		GetRentals(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		GetDefaultSettingsForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		UpdateDefaultSettingsForRental(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/rentals/{id}/variableSettings", func(w http.ResponseWriter, r *http.Request) {
		GetVariableSettingsForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/variableSettings", func(w http.ResponseWriter, r *http.Request) {
		UpdateVariableSettingsForRental(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/rentals/{id}/variableSettings", func(w http.ResponseWriter, r *http.Request) {
		CreateVariableSettingsForRental(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentals/{id}/variableSettings", func(w http.ResponseWriter, r *http.Request) {
		DeleteVariableSettingsForRental(w, r, db)
	}).Methods("DELETE")

	r.HandleFunc("/rentals/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		GetRentalTimeblocks(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalTimeblock(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentals/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		GetStatusForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		UpdateStatusForRental(w, r, db)
	}).Methods("PUT")

	//Boats
	r.HandleFunc("/boats", func(w http.ResponseWriter, r *http.Request) {
		GetBoats(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		GetDefaultSettingsForBoat(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		UpdateDefaultSettingsForBoat(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/boats/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		GetStatusForBoat(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		UpdateStatusForBoat(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/boats/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		GetBoatTimeblocks(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		CreateBoatTimeblock(w, r, db)
	}).Methods("POST")

	//Booking Status
	r.HandleFunc("/bookingStatus", func(w http.ResponseWriter, r *http.Request) {
		GetBookingStatus(w, r, db)
	}).Methods("GET")

	//Refund Status
	r.HandleFunc("/refundStatus", func(w http.ResponseWriter, r *http.Request) {
		GetRefundStatus(w, r, db)
	}).Methods("GET")

	//Alcohol
	r.HandleFunc("/alcohol", func(w http.ResponseWriter, r *http.Request) {
		GetAlcohol(w, r, db)
	}).Methods("GET")

	//OrderableAlcohol

	r.HandleFunc("/orderableAlcohol", func(w http.ResponseWriter, r *http.Request) {
		GetOrderableAlcohol(w, r, db)
	}).Methods("GET")

	//Alcohol Types
	r.HandleFunc("/alcoholTypes", func(w http.ResponseWriter, r *http.Request) {
		GetAlcoholTypes(w, r, db)
	}).Methods("GET")

	//Alcohol Quantity Types
	r.HandleFunc("/alcoholQuantityTypes", func(w http.ResponseWriter, r *http.Request) {
		GetAlcoholQuantityTypes(w, r, db)
	}).Methods("GET")

	//Event Types
	r.HandleFunc("/eventTypes", func(w http.ResponseWriter, r *http.Request) {
		GetEventTypes(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/eventTypes", func(w http.ResponseWriter, r *http.Request) {
		CreateEventType(w, r, db)
	}).Methods("POST")

	//venues
	r.HandleFunc("/venues", func(w http.ResponseWriter, r *http.Request) {
		GetVenues(w, r, db)
	}).Methods("GET")

	//Booking Cost Types
	r.HandleFunc("/bookingCostTypes", func(w http.ResponseWriter, r *http.Request) {
		GetBookingCostTypes(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookingCostTypes", func(w http.ResponseWriter, r *http.Request) {
		CreateBookingCostType(w, r, db)
	}).Methods("POST")

	//Locations
	r.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		GetLocations(w, r, db)
	}).Methods("GET")

	//Payment Methods
	r.HandleFunc("/paymentMethods", func(w http.ResponseWriter, r *http.Request) {
		GetPaymentMethods(w, r, db)
	}).Methods("GET")

	//Venue Event Types
	r.HandleFunc("/venueEventTypes", func(w http.ResponseWriter, r *http.Request) {
		GetVenueEventTypes(w, r, db)
	}).Methods("GET")

	//Venue Event Type Default Settings
	r.HandleFunc("/venueEventTypes/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		GetDefaultSettingsForVenueEventType(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/venueEventTypes/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		UpdateDefaultSettingsForVenueEventType(w, r, db)
	}).Methods("PUT")

	//rental booking
	r.HandleFunc("/rentalBooking", func(w http.ResponseWriter, r *http.Request) {
		GetRentalBookings(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentalBooking", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalBooking(w, r, db)
	}).Methods("POST")

	//boat booking
	r.HandleFunc("/boatBooking", func(w http.ResponseWriter, r *http.Request) {
		GetBoatBookings(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boatBooking", func(w http.ResponseWriter, r *http.Request) {
		CreateBoatBooking(w, r, db)
	}).Methods("POST")

	//rentalTimeblock

	r.HandleFunc("/rentalTimeblock/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetRentalTimeblock(w, r, db)
	}).Methods("GET")

	//boatTimeblock

	r.HandleFunc("/boatTimeblock/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetBoatTimeblock(w, r, db)
	}).Methods("GET")

}
