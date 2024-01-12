package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Boat struct {
	ID        int
	Name      string
	Occupancy int
	Weight    int
}

type BoatDefaultSettings struct {
	ID                      int
	BoatId                  int
	DailyCost               float32
	MinimunBookingDuration  int
	AdvertiseAtAllLocations bool
	fileId                  int
}

type BoatStatus struct {
	ID                int
	BoatId            int
	IsClean           bool
	LowFuel           bool
	CurrentLocationID int
}

func GetBoats(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Get all the boats from the database.

	rows, err := db.Query("SELECT * FROM boat")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var boats []Boat

	for rows.Next() {
		var boat Boat
		if err := rows.Scan(&boat.ID, &boat.Name, &boat.Occupancy, &boat.Weight); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		boats = append(boats, boat)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boats)

}

func GetDefaultSettingsForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	rows, err := db.Query("SELECT * FROM boat_default_settings WHERE boat_id = ?", boatID)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var defaultSettings []BoatDefaultSettings

	for rows.Next() {
		var defaultSetting BoatDefaultSettings
		if err := rows.Scan(&defaultSetting.ID, &defaultSetting.BoatId, &defaultSetting.DailyCost, &defaultSetting.MinimunBookingDuration, &defaultSetting.AdvertiseAtAllLocations, &defaultSetting.fileId); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		defaultSettings = append(defaultSettings, defaultSetting)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(defaultSettings)

}

func GetStatusForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	rows, err := db.Query("SELECT * FROM boat_status WHERE boat_id = ?", boatID)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var status []BoatStatus

	for rows.Next() {
		var boatStatus BoatStatus
		if err := rows.Scan(&boatStatus.ID, &boatStatus.BoatId, &boatStatus.IsClean, &boatStatus.LowFuel, &boatStatus.CurrentLocationID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		status = append(status, boatStatus)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)

}
