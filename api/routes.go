package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes for the API.
func InitRoutes(r *mux.Router, db *sql.DB) {
	// Define the routes.

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

	r.HandleFunc("/rentals/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		GetRentalTimeblocks(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/rentals/{id}/timeblocks", func(w http.ResponseWriter, r *http.Request) {
		CreateRentalTimeblock(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/rentals/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		GetStatusForRental(w, r, db)
	}).Methods("GET")

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

}
