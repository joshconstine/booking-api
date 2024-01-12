package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type VenueEventType struct {
	ID          int
	VenueID     int
	EventTypeID int
}

func GetVenueEventTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM venue_event_type")

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
