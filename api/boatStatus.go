package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BoatStatus struct {
	ID                int
	BoatId            int
	IsClean           bool
	LowFuel           bool
	CurrentLocationID int
}

func GetStatusForBoatId(boatId string, db *sql.DB) (BoatStatus, error) {

	rows, err := db.Query("SELECT * FROM boat_status WHERE boat_id = ?", boatId)

	if err != nil {
		return BoatStatus{}, err
	}

	defer rows.Close()

	var boatStatus BoatStatus

	if rows.Next() {
		err := rows.Scan(&boatStatus.ID, &boatStatus.BoatId, &boatStatus.IsClean, &boatStatus.LowFuel, &boatStatus.CurrentLocationID)
		if err != nil {
			return BoatStatus{}, err
		}
	}

	return boatStatus, nil
}

func GetStatusForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	boatStatus, err := GetStatusForBoatId(boatID, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatStatus)

}
func UpdateStatusForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	// Create a single instance of RentalUnitDefaultSettings.
	var boatStatus BoatStatus

	// Decode the JSON data.
	err := json.NewDecoder(r.Body).Decode(&boatStatus)
	if err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	// Update the database.
	updateQuery := "UPDATE boat_status SET is_clean = ?, low_fuel = ?, current_location_id = ? WHERE boat_id = ?"
	_, err = db.Exec(updateQuery, boatStatus.IsClean, boatStatus.LowFuel, boatStatus.CurrentLocationID, boatID)
	if err != nil {
		log.Fatalf("failed to update: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatStatus)

}
