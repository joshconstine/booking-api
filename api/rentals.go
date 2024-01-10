package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)
type Rental struct {
	ID int
	Name string
	LocationID int
	Bedrooms int
	Bathrooms int
}

type RentalWithLocation struct {
	Rental
	LocationName string

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

func GetRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the rental ID from the request URL.
	// We're using the gorilla/mux package to get the ID.
	//
	// For example, if the request URL is "/rentals/1",

	//get rental join location name from lcoation table

	rows, err := db.Query("SELECT rentals.id, rentals.name, locations.name, rentals.location_id, rentals.bathrooms, rentals.bedrooms FROM rentals JOIN locations ON rentals.location_id = locations.id WHERE rentals.id = ?", id)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var rentalWithLocation RentalWithLocation

	for rows.Next() {
		if err := rows.Scan(&rentalWithLocation.ID, &rentalWithLocation.Name, &rentalWithLocation.LocationName, &rentalWithLocation.LocationID, &rentalWithLocation.Bathrooms, &rentalWithLocation.Bedrooms); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalWithLocation)
}
