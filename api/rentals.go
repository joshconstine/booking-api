package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Rental struct {
	ID          int
	Name        string
	LocationID  int
	Bedrooms    int
	Bathrooms   int
	Description string
}

type RentalWithLocation struct {
	Rental
	LocationName string
}

type RentalInformtion struct {
	RentalID      int
	Name          string
	Description   string
	LocationID    int
	LocationName  string
	RentalIsClean bool
	Thumbnail     string
	Bookings      []RentalBookingDetails
	Timeblocks    []RentalTimeblock
}

func GetSingleRentalByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the rental ID from the request URL.
	// We're using the gorilla/mux package to get the ID.
	//
	// For example, if the request URL is "/rentals/1",

	//get rental join location name from lcoation table

	rows, err := db.Query("SELECT rental.id, rental.name, location.name, rental.location_id, rental.bathrooms, rental.bedrooms, rental.description FROM rental JOIN location ON rental.location_id = location.id WHERE rental.id = ?", id)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var rentalWithLocation RentalWithLocation

	for rows.Next() {
		if err := rows.Scan(&rentalWithLocation.ID, &rentalWithLocation.Name, &rentalWithLocation.LocationName, &rentalWithLocation.LocationID, &rentalWithLocation.Bathrooms, &rentalWithLocation.Bedrooms, &rentalWithLocation.Description); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalWithLocation)
}

func GetRentals(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database for all rentals.
	rows, err := db.Query("SELECT id, name, location_id, bedrooms, bathrooms, description FROM rental")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rentals to hold the data.
	var rentals []Rental

	// Loop through the data and insert into the rentals slice.
	for rows.Next() {
		var rental Rental
		if err := rows.Scan(&rental.ID, &rental.Name, &rental.LocationID, &rental.Bedrooms, &rental.Bathrooms, &rental.Description); err != nil {
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

	rows, err := db.Query("SELECT rental.id, rental.name, location.name, rental.location_id, rental.bathrooms, rental.bedrooms, rental.description FROM rental JOIN location ON rental.location_id = location.id WHERE rental.id = ?", id)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var rentalWithLocation RentalWithLocation

	for rows.Next() {
		if err := rows.Scan(&rentalWithLocation.ID, &rentalWithLocation.Name, &rentalWithLocation.LocationName, &rentalWithLocation.LocationID, &rentalWithLocation.Bathrooms, &rentalWithLocation.Bedrooms, &rentalWithLocation.Description); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalWithLocation)
}

func GetRentalInformation(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Get the rental ID from the request URL.
	// We're using the gorilla/mux package to get the ID.
	//
	// For example, if the request URL is "/rentals/1",

	//get rental join location name from lcoation table

	rows, err := db.Query("SELECT rental.id, rental.name, location.name, rental.location_id, rental.description, rs.is_clean FROM rental JOIN location ON rental.location_id = location.id JOIN rental_status rs ON rental.id = rs.rental_id")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var rentalInformation []RentalInformtion

	for rows.Next() {
		var rentalInfo RentalInformtion
		if err := rows.Scan(&rentalInfo.RentalID, &rentalInfo.Name, &rentalInfo.LocationName, &rentalInfo.LocationID, &rentalInfo.Description, &rentalInfo.RentalIsClean); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		threeMonthsFromNow := time.Now().AddDate(0, 3, 0)
		today := time.Now()

		rentalTimeblocks, err := GetRentalTimeblocksByRentalIDForRange(rentalInfo.RentalID, today, threeMonthsFromNow, db)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		rentalBookings, err := GetRentalBookingDetailsByRentalIdForRange(rentalInfo.RentalID, today, threeMonthsFromNow, db)

		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		thumbnail, err := GetRentalThumbnailByRentalID(rentalInfo.RentalID, db)

		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		rentalInfo.Thumbnail = thumbnail
		rentalInfo.Bookings = rentalBookings
		rentalInfo.Timeblocks = rentalTimeblocks

		rentalInformation = append(rentalInformation, rentalInfo)

	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalInformation)
}

func GetRentalThumbnailByRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	idStr, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("failed to convert id to int: %v", err)
	}
	thumbnail, err := GetRentalThumbnailByRentalID(idStr, db)
	if err != nil {
		log.Fatalf("failed to get thumbnail: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thumbnail)
}
