package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type BoatTimeblock struct {
	ID            int
	BoatID        int
	StartTime     time.Time
	EndTime       time.Time
	BoatBookingID *int
}

// Attempt to insert a boat timeblock into the database
func AttemptToInsertBoatTimeblock(db *sql.DB, boatID string, startTime time.Time, endTime time.Time, boatBookingID *int) (int, error) {
	// Format time values as strings in the MySQL datetime format.
	startTimeStr := startTime.Format("2006-01-02 15:04:05")
	endTimeStr := endTime.Format("2006-01-02 15:04:05")

	// Check for overlapping timeblocks
	overlapQuery := "SELECT id FROM boat_timeblock WHERE boat_id = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?))"
	rows, err := db.Query(overlapQuery, boatID, startTimeStr, startTimeStr, endTimeStr, endTimeStr, startTimeStr, endTimeStr)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	// If there are overlapping timeblocks, return false
	if rows.Next() {
		return -1, nil
	}

	// Insert the data into the database
	result, err := db.Exec("INSERT INTO boat_timeblock (boat_id, start_time, end_time, boat_booking_id) VALUES (?, ?, ?, ?)", boatID, startTimeStr, endTimeStr, boatBookingID)

	// Check if the error is a duplicate entry error
	if IsDuplicateKeyError(err) {
		// Handle duplicate entry error
		return -1, nil
	} else if err != nil {
		// Handle other errors
		return -1, err
	}

	// boat timeblock was successfully created

	// Get the ID of the newly created boat timeblock
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func GetBoatTimeblockById(id string, db *sql.DB) (BoatTimeblock, error) {
	// Query the database for the boat timeblock of the id.
	query := "SELECT * FROM boat_timeblock WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return BoatTimeblock{}, err
	}
	defer rows.Close()

	// Create a single instance of boatTimeblock.
	var boatTimeblock BoatTimeblock

	// Check if there is at least one row.
	if rows.Next() {
		var startTimeStr, endTimeStr string
		var boatBookingID *int

		// Scan the values into variables.
		if err := rows.Scan(&boatTimeblock.ID, &boatTimeblock.BoatID, &startTimeStr, &endTimeStr, &boatBookingID); err != nil {
			return BoatTimeblock{}, err
		}

		// Convert the datetime strings to time.Time.
		boatTimeblock.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return BoatTimeblock{}, err
		}

		boatTimeblock.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			return BoatTimeblock{}, err
		}

		boatTimeblock.BoatBookingID = boatBookingID
	}

	return boatTimeblock, nil
}

func RemoveBoatTimeblockById(id string, db *sql.DB) error {
	// Query the database for the boat timeblock of the id.
	query := "DELETE FROM boat_timeblock WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetBoatTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	timeblock, err := GetBoatTimeblockById(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblock)

}

func GetBoatTimeblocks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database.
	rows, err := db.Query("SELECT * FROM boat_timeblock WHERE boat_id = ?", id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of boats to hold the data.
	var timeblocks []BoatTimeblock

	// Iterate over the rows, adding the boats to the slice.
	for rows.Next() {
		var timeblock BoatTimeblock
		var startTimeStr, endTimeStr string
		var boatBookingID *int

		// Scan the values into variables.
		if err := rows.Scan(&timeblock.ID, &timeblock.BoatID, &startTimeStr, &endTimeStr, &boatBookingID); err != nil {
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

		timeblock.BoatBookingID = boatBookingID

		timeblocks = append(timeblocks, timeblock)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblocks)

}

func CreateBoatTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Parse the request body.
	var timeblock BoatTimeblock
	if err := json.NewDecoder(r.Body).Decode(&timeblock); err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	// Insert the data into the database.
	createdId, err := AttemptToInsertBoatTimeblock(db, id, timeblock.StartTime, timeblock.EndTime, timeblock.BoatBookingID)
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
func RemoveBoatTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := RemoveBoatTimeblockById(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	w.WriteHeader(http.StatusNoContent)
}
