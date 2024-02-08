package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RentalAmenity struct {
	ID      int
	Amenity Amenity
}

type CreateRentalAmenityRequest struct {
	AmenityID int
	RentalID  int
}

func DeleteRentalAmenity(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Delete a rental amenity by its ID.

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM rental_amenity WHERE id = ?", id)

	if err != nil {
		log.Fatalf("failed to delete rental amenity: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func CreateRentalAmenity(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Create a rental amenity.

	var request CreateRentalAmenityRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	_, err := db.Exec("INSERT INTO rental_amenity (amenity_id, rental_id) VALUES (?, ?)", request.AmenityID, request.RentalID)

	if err != nil {
		log.Fatalf("failed to insert rental amenity: %v", err)
	}

	w.WriteHeader(http.StatusCreated)

}

func GetAmenitiesForRentalID(id string, db *sql.DB) ([]RentalAmenity, error) {
	//
	// Get all the amenities for a rental by its ID.

	rows, err := db.Query("SELECT a.id, a.name, a.amenity_type_id, ra.id FROM amenity a JOIN rental_amenity ra ON a.id = ra.amenity_id WHERE ra.rental_id = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var amenities []RentalAmenity

	for rows.Next() {
		var amenity RentalAmenity
		if err := rows.Scan(&amenity.Amenity.ID, &amenity.Amenity.Name, &amenity.Amenity.AmenityTypeID, &amenity.ID); err != nil {
			return nil, err
		}
		amenities = append(amenities, amenity)
	}

	return amenities, nil

}

func GetAmenitiesForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all the amenities for a rental by its ID.

	vars := mux.Vars(r)
	rentalID := vars["id"]

	amenities, err := GetAmenitiesForRentalID(rentalID, db)

	if err != nil {
		log.Fatalf("failed to get amenities: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(amenities)
}
