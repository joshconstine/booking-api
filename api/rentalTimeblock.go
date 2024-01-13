package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type RentalTimeblock struct {
	ID              int
	RentalID        int
	StartTime       time.Time
	EndTime         time.Time
	RentalBookingID *int
}

// Attempt to insert a rental timeblock into the database
func attemptToInsertRentalTimeblock(db *sql.DB, rentalID string, startTimeStr string, endTimeStr string, rentalBookingID *int) (bool, error) {
	// Check for overlapping timeblocks
	overlapQuery := "SELECT id FROM rental_timeblock WHERE rental_id = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?))"
	rows, err := db.Query(overlapQuery, rentalID, startTimeStr, startTimeStr, endTimeStr, endTimeStr, startTimeStr, endTimeStr)
	if err != nil {
		return false, err

	}
	defer rows.Close()

	// If there are overlapping timeblocks, return false
	if rows.Next() {
		return false, nil
	}

	// Insert the data into the database
	_, err = db.Exec("INSERT INTO rental_timeblock (rental_id, start_time, end_time, rental_booking_id) VALUES (?, ?, ?, ?)", rentalID, startTimeStr, endTimeStr, rentalBookingID)

	// Check if the error is a duplicate entry error
	if IsDuplicateKeyError(err) {
		// Handle duplicate entry error
		return false, nil
	} else if err != nil {
		// Handle other errors
		return false, err
	}

	// Rental timeblock was successfully created
	return true, nil
}

func GetRentalTimeblocks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database.
	rows, err := db.Query("SELECT * FROM rental_timeblock WHERE rental_id = ?", id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rentals to hold the data.
	var timeblocks []RentalTimeblock

	// Iterate over the rows, adding the rentals to the slice.
	for rows.Next() {
		var timeblock RentalTimeblock
		var startTimeStr, endTimeStr string
		var rentalBookingID *int

		// Scan the values into variables.
		if err := rows.Scan(&timeblock.ID, &timeblock.RentalID, &startTimeStr, &endTimeStr, &rentalBookingID); err != nil {
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

		timeblock.RentalBookingID = rentalBookingID

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
	// Format time values as strings in the MySQL datetime format.
	startTimeStr := timeblock.StartTime.Format("2006-01-02 15:04:05")
	endTimeStr := timeblock.EndTime.Format("2006-01-02 15:04:05")

	// Insert the data into the database.
	created, err := attemptToInsertRentalTimeblock(db, id, startTimeStr, endTimeStr, timeblock.RentalBookingID)

	// Check for errors.
	if err != nil {
		log.Printf("failed to insert: %v", err)
		w.WriteHeader(http.StatusInternalServerError) // HTTP 500 Internal Server Error
		w.Write([]byte("Internal Server Error"))
		return
	}

	// Check if the timeblock was created.
	if !created {
		w.WriteHeader(http.StatusConflict) // HTTP 409 Conflict
		w.Write([]byte("Conflict: The timeblock already exists."))
		return
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	json.NewEncoder(w).Encode(timeblock)

}
