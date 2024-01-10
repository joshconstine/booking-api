package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)
type Rental struct {
	ID int
	Name string
	LocationID int
	Bedrooms int
	Bathrooms int
}


func GetRentals(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database for all rentals.
	rows, err := db.Query("SELECT * FROM rentals")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rentals to hold the data.
	var rentals []Rental

	// Loop through the data and insert into the rentals slice.
	for rows.Next() {
		var rental Rental
		if err := rows.Scan(&rental.ID, &rental.Name, &rental.LocationID, &rental.Bedrooms, &rental.Bathrooms); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		rentals = append(rentals, rental)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentals)
}