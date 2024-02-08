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
