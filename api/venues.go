package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Venue struct {
	ID         int
	Name       string
	LocationID int
}

func GetSingleVenueByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	venueID := vars["id"]

	// Query the database to get the venue
	rows, err := db.Query("SELECT venue.id, venue.name, venue.location_id FROM venue WHERE venue.id = ?", venueID)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var venue Venue

	for rows.Next() {
		if err := rows.Scan(&venue.ID, &venue.Name, &venue.LocationID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venue)
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
