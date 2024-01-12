package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RentalTimeblock struct {
	ID        int
	RentalID  int
	StartTime string
	EndTime   string
}

func GetRentalTimeblocks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	rows, err := db.Query("SELECT rental_timeblock.id, rental_timeblock.rental_id, rental_timeblock.start_time, rental_timeblock.end_time FROM rental_timeblock JOIN rentals ON rental_timeblock.rental_id = rentals.id WHERE rentals.id = ?", id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	// Create a slice of rentals to hold the data.
	var timeblocks []RentalTimeblock

	// Loop through the data and insert into the rentals slice.
	for rows.Next() {
		var timeblock RentalTimeblock
		if err := rows.Scan(&timeblock.ID, &timeblock.RentalID, &timeblock.StartTime, &timeblock.EndTime); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		timeblocks = append(timeblocks, timeblock)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblocks)
}

func CreateRentalTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Parse the request body.
	var timeblock RentalTimeblock
	if err := json.NewDecoder(r.Body).Decode(&timeblock); err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	// Insert the data into the database.
	_, err := db.Exec("INSERT INTO rental_timeblock (rental_id, start_time, end_time) VALUES (?, ?, ?)", id, timeblock.StartTime, timeblock.EndTime)
	if err != nil {
		log.Fatalf("failed to insert: %v", err)
	}

	// Return a status code.
	w.WriteHeader(http.StatusCreated)
}
