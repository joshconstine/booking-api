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
