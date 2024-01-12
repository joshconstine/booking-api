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

func GetDefaultSettingsForVenueEventType(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

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

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venueEventTypeDefaultSettings)
}
