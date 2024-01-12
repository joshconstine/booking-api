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

func GetStatusForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	rows, err := db.Query("SELECT * FROM boat_status WHERE boat_id = ?", boatID)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var boatStatus BoatStatus

	if rows.Next() {
		err := rows.Scan(&boatStatus.ID, &boatStatus.BoatId, &boatStatus.IsClean, &boatStatus.LowFuel, &boatStatus.CurrentLocationID)
		if err != nil {
			log.Fatalf("failed to scan: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatStatus)

}
