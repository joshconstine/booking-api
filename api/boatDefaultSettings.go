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
	BoatId                  int
	DailyCost               float32
	MinimunBookingDuration  int
	AdvertiseAtAllLocations bool
	fileId                  int
}

func GetDefaultSettingsForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	rows, err := db.Query("SELECT * FROM boat_default_settings WHERE boat_id = ?", boatID)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var defaultSettings BoatDefaultSettings

	if rows.Next() {
		err := rows.Scan(&defaultSettings.ID, &defaultSettings.BoatId, &defaultSettings.DailyCost, &defaultSettings.MinimunBookingDuration, &defaultSettings.AdvertiseAtAllLocations, &defaultSettings.fileId)
		if err != nil {
			log.Fatalf("failed to scan: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(defaultSettings)

}
