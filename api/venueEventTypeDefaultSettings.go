package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type VenueEventTypeDefaultSettings struct {
	ID                     int
	VenueEventTypeID       int
	HourlyRate             float32
	MinimumBookingDuration int
	FlatFee                float32
	EarliestBookingTime    string
	LatestBookingTime      string
}

func GetDefaultSettingsForVenueEventTypeID(id string, db *sql.DB) (VenueEventTypeDefaultSettings, error) {
	// Query the database for the default setting of the id.
	query := "SELECT * FROM venue_event_type_default_settings WHERE venue_event_type_id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a single instance of VenueEventTypeDefaultSettings.
	var venueEventTypeDefaultSettings VenueEventTypeDefaultSettings

	// Check if there is at least one row.
	if rows.Next() {
		err := rows.Scan(
			&venueEventTypeDefaultSettings.ID,
			&venueEventTypeDefaultSettings.VenueEventTypeID,
			&venueEventTypeDefaultSettings.HourlyRate,
			&venueEventTypeDefaultSettings.MinimumBookingDuration,
			&venueEventTypeDefaultSettings.FlatFee,
			&venueEventTypeDefaultSettings.EarliestBookingTime,
			&venueEventTypeDefaultSettings.LatestBookingTime,
		)
		if err != nil {
			log.Fatalf("failed to scan: %v", err)
		}
	}

	return venueEventTypeDefaultSettings, nil
}

func GetDefaultSettingsForVenueEventType(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the default settings for the venue event type.
	venueEventTypeDefaultSettings, err := GetDefaultSettingsForVenueEventTypeID(id, db)

	// Check if there is an error.
	if err != nil {
		log.Fatalf("failed to get default settings: %v", err)
	}
	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venueEventTypeDefaultSettings)
}

func UpdateDefaultSettingsForVenueEventType(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Create a single instance of VenueEventTypeDefaultSettings.
	var venueEventTypeDefaultSettings VenueEventTypeDefaultSettings

	// Decode the JSON data.
	err := json.NewDecoder(r.Body).Decode(&venueEventTypeDefaultSettings)
	if err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	// Update the database.
	query := "UPDATE venue_event_type_default_settings SET hourly_rate = ?, minimum_booking_duration = ?, flat_fee = ?, earliest_booking_time = ?, latest_booking_time = ? WHERE venue_event_type_id = ?"
	_, err = db.Exec(query, venueEventTypeDefaultSettings.HourlyRate, venueEventTypeDefaultSettings.MinimumBookingDuration, venueEventTypeDefaultSettings.FlatFee, venueEventTypeDefaultSettings.EarliestBookingTime, venueEventTypeDefaultSettings.LatestBookingTime, id)
	if err != nil {
		log.Fatalf("failed to update: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venueEventTypeDefaultSettings)
}
