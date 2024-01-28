package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BoatDefaultSettings struct {
	ID                      int
	BoatID                  int
	DailyCost               float64
	MinimumBookingDuration  int
	AdvertiseAtAllLocations bool
	FileID                  int
}

func GetDefaultSettingsForBoatId(boatId string, db *sql.DB) (BoatDefaultSettings, error) {

	rows, err := db.Query("SELECT * FROM boat_default_settings WHERE boat_id = ?", boatId)

	if err != nil {
		return BoatDefaultSettings{}, err
	}

	defer rows.Close()

	var defaultSettings BoatDefaultSettings

	if rows.Next() {
		err := rows.Scan(&defaultSettings.ID, &defaultSettings.BoatID, &defaultSettings.DailyCost, &defaultSettings.MinimumBookingDuration, &defaultSettings.AdvertiseAtAllLocations, &defaultSettings.FileID)
		if err != nil {
			return BoatDefaultSettings{}, err
		}
	}

	return defaultSettings, nil
}

func GetDefaultSettingsForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	defaultSettings, err := GetDefaultSettingsForBoatId(boatID, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(defaultSettings)

}
func UpdateDefaultSettingsForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	// Create a single instance of RentalUnitDefaultSettings.
	var defaultSettings BoatDefaultSettings

	// Decode the JSON data.
	err := json.NewDecoder(r.Body).Decode(&defaultSettings)
	if err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	// Update the database.
	query := "UPDATE boat_default_settings SET daily_cost = ?, minimum_booking_duration = ?, advertise_at_all_locations = ? WHERE boat_id = ?"
	_, err = db.Exec(query, defaultSettings.DailyCost, defaultSettings.MinimumBookingDuration, defaultSettings.AdvertiseAtAllLocations, boatID)
	if err != nil {
		log.Fatalf("failed to update: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(defaultSettings)
}
