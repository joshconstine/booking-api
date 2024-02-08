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
	ID      int
	BedType BedType
}

type RentalBedroomWithBeds struct {
	ID            int
	RentalID      int
	Name          string
	Description   string
	Floor         int
	RentalPhotoID *int
	Beds          []RentalBedroomBed
}

type CreateRentalBedroomBedRequest struct {
	RentalBedroomID int
	BedTypeID       int
}

func CreateRentalBedroomBed(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var request CreateRentalBedroomBedRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	_, err := db.Exec("INSERT INTO rental_bedroom_bed (rental_bedroom_id, bed_type_id) VALUES (?, ?)", request.RentalBedroomID, request.BedTypeID)

	if err != nil {
		log.Fatalf("failed to insert rental bedroom bed: %v", err)
	}

	w.WriteHeader(http.StatusCreated)

}
func DeleteRentalBedroomBed(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM rental_bedroom_bed WHERE id = ?", id)

	if err != nil {
		log.Fatalf("failed to delete rental bedroom bed: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateRentalBedroom(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var bedroom RentalBedroom
	if err := json.NewDecoder(r.Body).Decode(&bedroom); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	_, err := db.Exec("UPDATE rental_bedroom SET name = ?, description = ?, floor = ?, rental_photo_id = ? WHERE id = ?", bedroom.Name, bedroom.Description, bedroom.Floor, bedroom.RentalPhotoID, bedroom.ID)

	if err != nil {
		log.Fatalf("failed to update rental bedroom: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func CreateRentalBedroom(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var bedroom RentalBedroom
	if err := json.NewDecoder(r.Body).Decode(&bedroom); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	_, err := db.Exec("INSERT INTO rental_bedroom (rental_id, name, description, floor, rental_photo_id) VALUES (?, ?, ?, ?, ?)", bedroom.RentalID, bedroom.Name, bedroom.Description, bedroom.Floor, bedroom.RentalPhotoID)

	if err != nil {
		log.Fatalf("failed to insert rental bedroom: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteRentalBedroom(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM rental_bedroom WHERE id = ?", id)

	if err != nil {
		log.Fatalf("failed to delete rental bedroom: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func GetBedsForRentalBedroomID(id int, db *sql.DB) ([]RentalBedroomBed, error) {
	rows, err := db.Query("SELECT bed_type.id, bed_type.name, rental_bedroom_bed.id FROM bed_type JOIN rental_bedroom_bed ON bed_type.id = rental_bedroom_bed.bed_type_id WHERE rental_bedroom_bed.rental_bedroom_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beds []RentalBedroomBed
	for rows.Next() {
		var bed RentalBedroomBed
		err := rows.Scan(&bed.BedType.ID, &bed.BedType.Name, &bed.ID)
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
