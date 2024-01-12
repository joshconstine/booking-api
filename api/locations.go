package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Location struct {
	ID   int
	Name string
}

// GetLocations returns all locations from the database.
func GetLocations(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM location")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var locations []Location

	for rows.Next() {
		var location Location
		if err := rows.Scan(&location.ID, &location.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		locations = append(locations, location)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)

}
