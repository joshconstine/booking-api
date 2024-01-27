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

type RentalTimeblock struct {
	ID              int
	RentalID        int
	StartTime       time.Time
	EndTime         time.Time
	RentalBookingID *int
}

// Attempt to insert a rental timeblock into the database
func AttemptToInsertRentalTimeblock(db *sql.DB, rentalID string, startTime time.Time, endTime time.Time, rentalBookingID *int) (int, error) {
	// Format time values as strings in the MySQL datetime format.
	startTimeStr := startTime.Format("2006-01-02 15:04:05")
	endTimeStr := endTime.Format("2006-01-02 15:04:05")

	// Check for overlapping timeblocks
	overlapQuery := "SELECT id FROM rental_timeblock WHERE rental_id = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?))"
	rows, err := db.Query(overlapQuery, rentalID, startTimeStr, startTimeStr, endTimeStr, endTimeStr, startTimeStr, endTimeStr)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	// If there are overlapping timeblocks, return false
	if rows.Next() {
		return -1, nil
	}

	// Insert the data into the database
	result, err := db.Exec("INSERT INTO rental_timeblock (rental_id, start_time, end_time, rental_booking_id) VALUES (?, ?, ?, ?)", rentalID, startTimeStr, endTimeStr, rentalBookingID)

	// Check if the error is a duplicate entry error
	if IsDuplicateKeyError(err) {
		// Handle duplicate entry error
		return -1, nil
	} else if err != nil {
		// Handle other errors
		return -1, err
	}

	// Rental timeblock was successfully created

	// Get the ID of the newly created rental timeblock
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func GetRentalTimeblockById(id string, db *sql.DB) (RentalTimeblock, error) {
	// Query the database for the rental timeblock of the id.
	query := "SELECT * FROM rental_timeblock WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return RentalTimeblock{}, err
	}
	defer rows.Close()

	// Create a single instance of RentalTimeblock.
	var rentalTimeblock RentalTimeblock

	// Check if there is at least one row.
	if rows.Next() {
		var startTimeStr, endTimeStr string
		var rentalBookingID *int

		// Scan the values into variables.
		if err := rows.Scan(&rentalTimeblock.ID, &rentalTimeblock.RentalID, &startTimeStr, &endTimeStr, &rentalBookingID); err != nil {
			return RentalTimeblock{}, err
		}

		// Convert the datetime strings to time.Time.
		rentalTimeblock.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return RentalTimeblock{}, err
		}

		rentalTimeblock.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			return RentalTimeblock{}, err
		}

		rentalTimeblock.RentalBookingID = rentalBookingID
	}

	return rentalTimeblock, nil
}

func GetRentalTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	timeblock, err := GetRentalTimeblockById(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblock)

}

func GetRentalTimeblocksByRentalID(rentalID int, db *sql.DB) ([]RentalTimeblock, error) {
	rentalIdString := strconv.Itoa(rentalID)

	// Query the database for the rental timeblock of the id.
	query := "SELECT * FROM rental_timeblock WHERE rental_id = ?"
	rows, err := db.Query(query, rentalIdString)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		// Convert the datetime strings to time.Time.
		timeblock.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return nil, err
		}

		timeblock.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			return nil, err
		}

		timeblock.RentalBookingID = rentalBookingID

		timeblocks = append(timeblocks, timeblock)
	}

	return timeblocks, nil
}
func GetRentalTimeblocksByRentalIDForRange(rentalID int, from time.Time, to time.Time, db *sql.DB) ([]RentalTimeblock, error) {
	rentalIdString := strconv.Itoa(rentalID)
	fromString := from.Format("2006-01-02 15:04:05")
	toString := to.Format("2006-01-02 15:04:05")

	// Modify the query to filter by start and end time
	query := `SELECT * FROM rental_timeblock WHERE rental_id = ? AND start_time >= ? AND end_time <= ?`
	rows, err := db.Query(query, rentalIdString, fromString, toString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timeblocks []RentalTimeblock

	for rows.Next() {
		var timeblock RentalTimeblock
		var startTimeStr, endTimeStr string
		var rentalBookingID *int

		if err := rows.Scan(&timeblock.ID, &timeblock.RentalID, &startTimeStr, &endTimeStr, &rentalBookingID); err != nil {
			return nil, err
		}

		timeblock.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return nil, err
		}

		timeblock.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			return nil, err
		}

		timeblock.RentalBookingID = rentalBookingID

		timeblocks = append(timeblocks, timeblock)
	}

	return timeblocks, nil
}

func RemoveRentalTimeblockById(id string, db *sql.DB) error {

	_, err := db.Exec("DELETE FROM rental_timeblock WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func RemoveRentalTimeblock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := RemoveRentalTimeblockById(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

}

func GetRentalTimeblocks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	intId, err := strconv.Atoi(id)

	timeblocks, err := GetRentalTimeblocksByRentalID(intId, db)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
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
	createdId, err := AttemptToInsertRentalTimeblock(db, id, timeblock.StartTime, timeblock.EndTime, timeblock.RentalBookingID)
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
