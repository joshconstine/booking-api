package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes for the API.
func InitRoutes(r *mux.Router, db *sql.DB) {
	// Define the routes.

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

	//Refund Status
	r.HandleFunc("/refundStatus", func(w http.ResponseWriter, r *http.Request) {
		GetRefundStatus(w, r, db)
	}).Methods("GET")

	//Users
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		GetUsers(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	}).Methods("POST")
}
