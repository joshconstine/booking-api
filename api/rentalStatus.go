package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RentalStatus struct {
	ID       int
	RentalID int
	IsClean  bool
}

func GetStatusForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	rentalId := vars["id"]

	rows, err := db.Query("SELECT * FROM rental_status WHERE rental_id = ?", rentalId)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var rentalStatus []RentalStatus

	for rows.Next() {
		var status RentalStatus
		if err := rows.Scan(&status.ID, &status.RentalID, &status.IsClean); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		rentalStatus = append(rentalStatus, status)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalStatus)

}
