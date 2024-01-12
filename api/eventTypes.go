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

func CreateEventType(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var eventType EventType
	json.NewDecoder(r.Body).Decode(&eventType)

	// Insert the data into the database.
	_, err := db.Exec("INSERT INTO event_type (name) VALUES (?)", eventType.Name)

	//checkfor Duplicate entry
	if err != nil {
		// Check if the error is a duplicate entry error
		if IsDuplicateKeyError(err) {
			// Handle duplicate entry error
			w.WriteHeader(http.StatusConflict) // HTTP 409 Conflict
			w.Write([]byte("Duplicate entry: The booking cost type already exists."))
		} else {
			// Handle other errors
			log.Printf("failed to insert: %v", err)
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500 Internal Server Error
			w.Write([]byte("Internal Server Error"))
		}
		return
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventType)

}
