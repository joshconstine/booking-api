package api

import (
	"database/sql"
	"log"
	"net/http"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type RentalBathroom struct {
	ID            int
	RentalID      int
	Name          string
	Description   string
	Floor         int
	Bathtub       bool
	Shower        bool
	RentalPhotoID *int
}

func GetBathroomsForRentalID(id string, db *sql.DB) ([]RentalBathroom, error) {
	rows, err := db.Query("SELECT id, rental_id, name, description, floor, bathtub, shower, rental_photo_id FROM rental_bathroom WHERE rental_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bathrooms []RentalBathroom
	for rows.Next() {
		var bathroom RentalBathroom
		err := rows.Scan(&bathroom.ID, &bathroom.RentalID, &bathroom.Name, &bathroom.Description, &bathroom.Floor, &bathroom.Bathtub, &bathroom.Shower, &bathroom.RentalPhotoID)
		if err != nil {
			return nil, err
		}
		bathrooms = append(bathrooms, bathroom)
	}
	return bathrooms, nil
}

func GetBathroomsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	rentalID := vars["id"]

	bathrooms, err := GetBathroomsForRentalID(rentalID, db)
	if err != nil {
		log.Fatalf("failed to get bathrooms: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bathrooms)
}
