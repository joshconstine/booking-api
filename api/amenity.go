package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Amenity struct {
	ID            int
	Name          string
	AmenityTypeID int
}

type AmenityType struct {
	ID   int
	Name string
}

type AmenityWithTypeName struct {
	ID       int
	Name     string
	TypeName string
}

func GetAmenityTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all the amenity types from the database.

	rows, err := db.Query("SELECT id, name FROM amenity_type")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var amenityTypes []AmenityType

	for rows.Next() {
		var amenityType AmenityType
		if err := rows.Scan(&amenityType.ID, &amenityType.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		amenityTypes = append(amenityTypes, amenityType)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(amenityTypes)
}

func GetAmenities(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all the amenities from the database wiht their type name.

	rows, err := db.Query("SELECT a.id, a.name, at.name FROM amenity a JOIN amenity_type at ON a.amenity_type_id = at.id")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var amenities []AmenityWithTypeName

	for rows.Next() {
		var amenity AmenityWithTypeName
		if err := rows.Scan(&amenity.ID, &amenity.Name, &amenity.TypeName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		amenities = append(amenities, amenity)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(amenities)

}
