package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type VenueEventType struct {
	ID          int
	VenueID     int
	EventTypeID int
}

func GetVenueEventTypesForSingleVenue(id string, db *sql.DB) ([]VenueEventType, error) {
	rows, err := db.Query("SELECT id, venue_id, event_type_id FROM venue_event_type WHERE venue_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var venueEventTypes []VenueEventType

	for rows.Next() {
		var eventType VenueEventType
		if err := rows.Scan(&eventType.ID, &eventType.VenueID, &eventType.EventTypeID); err != nil {
			return nil, err
		}
		venueEventTypes = append(venueEventTypes, eventType)
	}

	return venueEventTypes, nil
}

func GetVenueEventTypesForVenue(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	venueEventTypes, err := GetVenueEventTypesForSingleVenue(id, db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venueEventTypes)
}

func GetVenueEventTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT  id, venue_id, event_type_id FROM venue_event_type")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var venueEventTypes []VenueEventType

	for rows.Next() {
		var venueEventType VenueEventType
		if err := rows.Scan(&venueEventType.ID, &venueEventType.VenueID, &venueEventType.EventTypeID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		venueEventTypes = append(venueEventTypes, venueEventType)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venueEventTypes)

}
