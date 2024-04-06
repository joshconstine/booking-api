package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes for the API.
func InitRoutes(r *mux.Router, db *sql.DB) {
	// Define the routes.

	//Invoicing

	r.HandleFunc("/invoice", func(w http.ResponseWriter, r *http.Request) {
		CreateInvoiceHandler(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		CreateBooking(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/bookings/{id}/details", func(w http.ResponseWriter, r *http.Request) {
		GetBookingDetails(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings/{id}/invoice", func(w http.ResponseWriter, r *http.Request) {
		HandleCreateInvoiceForBookingHandler(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/bookings/{id}/details", func(w http.ResponseWriter, r *http.Request) {
		UpdateBookingDetailsForBooking(w, r, db)
	}).Methods("PUT")

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

	r.HandleFunc("/bookings/{id}/info", func(w http.ResponseWriter, r *http.Request) {
		GetBookingInformation(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings/{id}/invoice", func(w http.ResponseWriter, r *http.Request) {
		HandleGetInvoiceForBookingId(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/bookings/snapshots", func(w http.ResponseWriter, r *http.Request) {
		GetBookingSnapshots(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/info", func(w http.ResponseWriter, r *http.Request) {
		GetRentalInformation(w, r, db)
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

	r.HandleFunc("/variableSettings/{id}", func(w http.ResponseWriter, r *http.Request) {
		HandleDeleteVariableSettingsById(w, r, db)
	}).Methods("DELETE")

	r.HandleFunc("/rentals/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		GetRentalTimeblocks(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalTimeblock(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentals/{id}/bookings", func(w http.ResponseWriter, r *http.Request) {
		GetRentalBookingsForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		GetStatusForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		UpdateStatusForRental(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/rentals/{id}/photos/{photoID}", func(w http.ResponseWriter, r *http.Request) {
		DeleteRentalPhoto(w, r, db)
	}).Methods("DELETE")

	r.HandleFunc("/rentals/{id}/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		GetRentalThumbnailByRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/bedrooms", func(w http.ResponseWriter, r *http.Request) {
		GetBedroomsForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/bathrooms", func(w http.ResponseWriter, r *http.Request) {
		GetBathroomsForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/amenities", func(w http.ResponseWriter, r *http.Request) {
		GetAmenitiesForRental(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		GetDefaultSettingsForBoat(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/defaultSettings", func(w http.ResponseWriter, r *http.Request) {
		UpdateDefaultSettingsForBoat(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/boats/{id}/variableSettings", func(w http.ResponseWriter, r *http.Request) {
		GetVariableSettingsForBoat(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/variableSettings", func(w http.ResponseWriter, r *http.Request) {
		CreateVariableSettingsForBoat(w, r, db)
	}).Methods("POST")

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

	r.HandleFunc("/boats/{id}/bookings", func(w http.ResponseWriter, r *http.Request) {
		GetBoatBookingsForBoat(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boats/{id}/photos/{photoID}", func(w http.ResponseWriter, r *http.Request) {
		DeleteBoatPhoto(w, r, db)
	}).Methods("DELETE")

	r.HandleFunc("/boats/{id}/photos", func(w http.ResponseWriter, r *http.Request) {
		CreateBoatPhoto(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/boats/{id}/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		GetBoatThumbnail(w, r, db)
	}).Methods("GET")

	//Booking Status
	r.HandleFunc("/bookingStatus", func(w http.ResponseWriter, r *http.Request) {
		GetBookingStatus(w, r, db)
	}).Methods("GET")

	//Refund Status
	r.HandleFunc("/refundStatus", func(w http.ResponseWriter, r *http.Request) {
		GetRefundStatus(w, r, db)
	}).Methods("GET")

	//Booking Cost Types
	r.HandleFunc("/bookingCostTypes", func(w http.ResponseWriter, r *http.Request) {
		CreateBookingCostType(w, r, db)
	}).Methods("POST")

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

	r.HandleFunc("/rentalTimeblock/{id}", func(w http.ResponseWriter, r *http.Request) {
		RemoveRentalTimeblock(w, r, db)
	}).Methods("DELETE")

	//boatTimeblock

	r.HandleFunc("/boatTimeblock/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetBoatTimeblock(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/boatTimeblock/{id}", func(w http.ResponseWriter, r *http.Request) {
		RemoveBoatTimeblock(w, r, db)
	}).Methods("DELETE")

	//venueTimeblock
	r.HandleFunc("/venueTimeblock/{id}", func(w http.ResponseWriter, r *http.Request) {
		RemoveVenueTimeblock(w, r, db)
	}).Methods("DELETE")
	//Users
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		GetUsers(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	}).Methods("POST")

	//Booking Payments

	r.HandleFunc("/rentalAmenity", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalAmenity(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentalAmenity/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteRentalAmenity(w, r, db)
	}).Methods("DELETE")

	//rental Bedroom beds

	r.HandleFunc("/rentalBedroomBed", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalBedroomBed(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentalBedroomBed/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteRentalBedroomBed(w, r, db)
	}).Methods("DELETE")

	//rental Bedroom

	r.HandleFunc("/rentalBedroom", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalBedroom(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentalBedroom", func(w http.ResponseWriter, r *http.Request) {
		UpdateRentalBedroom(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/rentalBedroom/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteRentalBedroom(w, r, db)
	}).Methods("DELETE")

	//rental Bathroom

	r.HandleFunc("/rentalBathroom", func(w http.ResponseWriter, r *http.Request) {
		UpdateRentalBathroom(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/rentalBathroom", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalBathroom(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentalBathroom/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteRentalBathroom(w, r, db)
	}).Methods("DELETE")

}
