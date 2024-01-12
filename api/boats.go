package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Boat struct {
	ID        int
	Name      string
	Occupancy int
	Weight    int
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
