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

type AmenityWithType struct {
	ID            int
	Name          string
	TypeName      string
	AmenityTypeID int
}

func CreateAmenity(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Create an amenity.

	var amenity Amenity
	if err := json.NewDecoder(r.Body).Decode(&amenity); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	_, err := db.Exec("INSERT INTO amenity (name, amenity_type_id) VALUES (?, ?)", amenity.Name, amenity.AmenityTypeID)

	if err != nil {
		log.Fatalf("failed to insert amenity: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
}

func CreateAmenityType(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Create an amenity type.

	var amenityType AmenityType
	if err := json.NewDecoder(r.Body).Decode(&amenityType); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	_, err := db.Exec("INSERT INTO amenity_type (name) VALUES (?)", amenityType.Name)

	if err != nil {
		log.Fatalf("failed to insert amenity type: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
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

	rows, err := db.Query("SELECT a.id, a.name, a.amenity_type_id ,at.name FROM amenity a JOIN amenity_type at ON a.amenity_type_id = at.id")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var amenities []AmenityWithType

	for rows.Next() {
		var amenity AmenityWithType
		if err := rows.Scan(&amenity.ID, &amenity.Name, &amenity.AmenityTypeID, &amenity.TypeName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		amenities = append(amenities, amenity)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(amenities)

}
