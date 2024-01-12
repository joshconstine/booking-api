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
	rentalID := vars["id"]

	// Query the database for the status of the rental.
	query := "SELECT * FROM rental_status WHERE rental_id = ?"
	rows, err := db.Query(query, rentalID)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a single instance of RentalStatus.
	var rentalStatus RentalStatus

	// Check if there is at least one row.
	if rows.Next() {
		err := rows.Scan(&rentalStatus.ID, &rentalStatus.RentalID, &rentalStatus.IsClean)
		if err != nil {
			log.Fatalf("failed to scan: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalStatus)
}
