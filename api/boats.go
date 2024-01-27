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
	MaxWeight int
}

func GetBoats(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Get all the boats from the database.

	rows, err := db.Query("SELECT id, name, occupancy, max_weight FROM boat")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var boats []Boat

	for rows.Next() {
		var boat Boat
		if err := rows.Scan(&boat.ID, &boat.Name, &boat.Occupancy, &boat.MaxWeight); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		boats = append(boats, boat)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boats)
}
func GetSingleBoatByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the boat ID from the request URL.
	// We're using the gorilla/mux package to get the ID.
	//
	// For example, if the request URL is "/boats/1",

	//get boat join location name from lcoation table

	rows, err := db.Query("SELECT id, name, occupancy, max_weight FROM boat WHERE id = ?", id)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var boat Boat

	for rows.Next() {
		if err := rows.Scan(&boat.ID, &boat.Name, &boat.Occupancy, &boat.MaxWeight); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boat)

}
