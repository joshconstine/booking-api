package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type VenueTimeblock struct {
	ID        int
	VenueID   int
	StartTime time.Time
	EndTime   time.Time
	Note      *string
	EventID   *int
}

// Attempt to insert a venue timeblock into the database
func AttemptToInsertVenueTimeblock(db *sql.DB, venueID int, startTime time.Time, endTime time.Time, note *string, eventID *int) (int, error) {
	// Format time values as strings in the MySQL datetime format.
	startTimeStr := startTime.Format("2006-01-02 15:04:05")
	endTimeStr := endTime.Format("2006-01-02 15:04:05")

	// Check for overlapping timeblocks
	overlapQuery := "SELECT id FROM venue_timeblock WHERE venue_id = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?))"
	rows, err := db.Query(overlapQuery, venueID, startTimeStr, startTimeStr, endTimeStr, endTimeStr, startTimeStr, endTimeStr)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	// If there are overlapping timeblocks, return false
	if rows.Next() {
		return -1, nil
	}

	// Insert the data into the database
	result, err := db.Exec("INSERT INTO venue_timeblock (venue_id, start_time, end_time, note, event_id) VALUES (?, ?, ?, ?, ?)", venueID, startTimeStr, endTimeStr, note, eventID)

	// Check if the error is a duplicate entry error
	if IsDuplicateKeyError(err) {
		// Handle duplicate entry error
		return -1, nil
	} else if err != nil {
		// Handle other errors
		return -1, err
	}

	// venue timeblock was successfully created

	// Get the ID of the newly created venue timeblock
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func GetVenueTimeblockById(id string, db *sql.DB) (VenueTimeblock, error) {
	// Query the database for the venue timeblock of the id.
	query := "SELECT id, venue_id, start_time, end_time, note, event_id FROM venue_timeblock WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return VenueTimeblock{}, err
	}
	defer rows.Close()

	// Create a single instance of venueTimeblock.
	var venueTimeblock VenueTimeblock

	// Check if there is at least one row.
	if rows.Next() {
		var startTimeStr, endTimeStr string
		var eventID *int

		// Scan the values into variables.
		if err := rows.Scan(&venueTimeblock.ID, &venueTimeblock.VenueID, &startTimeStr, &endTimeStr, &venueTimeblock.Note, &eventID); err != nil {
			return VenueTimeblock{}, err
		}

		// Convert the datetime strings to time.Time.
		venueTimeblock.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return VenueTimeblock{}, err
		}

		venueTimeblock.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			return VenueTimeblock{}, err
		}

		venueTimeblock.EventID = eventID
	}

	return venueTimeblock, nil
}

func RemoveVenueTimeblockById(id string, db *sql.DB) error {
	// Query the database for the venue timeblock of the id.
	query := "DELETE FROM venue_timeblock WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetVenueTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	timeblock, err := GetVenueTimeblockById(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblock)

}

func GetVenueTimeblocks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database.
	rows, err := db.Query("SELECT  id, venue_id, start_time, end_time, note, event_id FROM venue_timeblock WHERE venue_id = ?", id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of venues to hold the data.
	var timeblocks []VenueTimeblock

	// Iterate over the rows, adding the venues to the slice.
	for rows.Next() {
		var timeblock VenueTimeblock
		var startTimeStr, endTimeStr string
		var eventID *int

		// Scan the values into variables.
		if err := rows.Scan(&timeblock.ID, &timeblock.VenueID, &startTimeStr, &endTimeStr, &timeblock.Note, &eventID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		// Convert the datetime strings to time.Time.
		timeblock.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			log.Fatalf("failed to parse start time: %v", err)
		}

		timeblock.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			log.Fatalf("failed to parse end time: %v", err)
		}

		timeblock.EventID = eventID

		timeblocks = append(timeblocks, timeblock)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblocks)

}

func CreateVenueTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Parse the request body.
	var timeblock VenueTimeblock
	if err := json.NewDecoder(r.Body).Decode(&timeblock); err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	idInt, err := strconv.Atoi(id)
	// Insert the data into the database.
	createdId, err := AttemptToInsertVenueTimeblock(db, idInt, timeblock.StartTime, timeblock.EndTime, timeblock.Note, timeblock.EventID)
	// Check for errors.
	if err != nil {
		log.Printf("failed to insert: %v", err)
		w.WriteHeader(http.StatusInternalServerError) // HTTP 500 Internal Server Error
		w.Write([]byte("Internal Server Error"))
		return
	}

	// Check if the timeblock was created.
	if createdId == -1 {
		w.WriteHeader(http.StatusConflict) // HTTP 409 Conflict
		w.Write([]byte("Conflict: The timeblock already exists."))
		return
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	json.NewEncoder(w).Encode(timeblock)

}
func RemoveVenueTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := RemoveVenueTimeblockById(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	w.WriteHeader(http.StatusNoContent)
}
