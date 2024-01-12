package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Venue struct {
	ID         int
	Name       string
	LocationID int
}

func GetVenues(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM venue")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var venues []Venue

	for rows.Next() {
		var venue Venue
		if err := rows.Scan(&venue.ID, &venue.Name, &venue.LocationID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		venues = append(venues, venue)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venues)

}
