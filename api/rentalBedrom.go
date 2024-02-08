package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type RentalBedroom struct {
	ID            int
	RentalID      int
	Name          string
	Description   string
	Floor         int
	RentalPhotoID *int
}

type RentalBedroomBed struct {
	ID              int
	RentalBedroomID int
	BedTypeID       int
}

type BedType struct {
	ID   int
	Name string
}

type RentalBedroomWithBeds struct {
	ID            int
	RentalID      int
	Name          string
	Description   string
	Floor         int
	RentalPhotoID *int
	Beds          []BedType
}

func GetBedsForRentalBedroomID(id int, db *sql.DB) ([]BedType, error) {
	rows, err := db.Query("SELECT bed_type.id, bed_type.name FROM bed_type JOIN rental_bedroom_bed ON bed_type.id = rental_bedroom_bed.bed_type_id WHERE rental_bedroom_bed.rental_bedroom_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beds []BedType
	for rows.Next() {
		var bed BedType
		err := rows.Scan(&bed.ID, &bed.Name)
		if err != nil {
			return nil, err
		}

		beds = append(beds, bed)
	}
	return beds, nil
}

func GetBedroomsForRentalIDWithBeds(id string, db *sql.DB) ([]RentalBedroomWithBeds, error) {

	rows, err := db.Query("SELECT id, rental_id, name, description, floor, rental_photo_id FROM rental_bedroom WHERE rental_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bedrooms []RentalBedroomWithBeds
	for rows.Next() {
		var bedroom RentalBedroomWithBeds
		err := rows.Scan(&bedroom.ID, &bedroom.RentalID, &bedroom.Name, &bedroom.Description, &bedroom.Floor, &bedroom.RentalPhotoID)
		if err != nil {
			return nil, err
		}

		beds, err := GetBedsForRentalBedroomID(bedroom.ID, db)
		if err != nil {
			return nil, err
		}

		bedroom.Beds = beds
		bedrooms = append(bedrooms, bedroom)

	}
	return bedrooms, nil
}

func GetBedroomsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	rentalID := vars["id"]

	bedrooms, err := GetBedroomsForRentalIDWithBeds(rentalID, db)
	if err != nil {
		log.Fatalf("failed to get bedrooms: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bedrooms)

}
