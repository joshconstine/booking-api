package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type EventType struct {
	ID   int
	Name string
}

func GetEventTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM event_type")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var eventTypes []EventType

	for rows.Next() {
		var eventType EventType
		if err := rows.Scan(&eventType.ID, &eventType.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		eventTypes = append(eventTypes, eventType)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventTypes)

}
