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

type BoatBooking struct {
	ID              int
	BoatID          int
	BookingID       int
	BoatTimeBlockID int
	BookingStatusID int
	BookingFileID   int
	LocationID      int
}

type RequestBoatBooking struct {
	BoatID     int
	BookingID  int
	StartTime  time.Time
	EndTime    time.Time
	LocationID int
}

type BoatBookingDetails struct {
	ID              int
	BoatID          int
	BookingID       int
	BoatTimeBlockID int
	BookingStatusID int
	BookingFileID   int
	LocationID      int
	StartTime       time.Time
	EndTime         time.Time
}

func AttemptToBookBoat(details RequestBoatBooking, db *sql.DB) (int64, error) {

	//oprn transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	boatIdString := strconv.Itoa(details.BoatID)

	// Attempt to create boat timeblock
	boatTimeblockID, err := AttemptToInsertBoatTimeblock(db, boatIdString, details.StartTime, details.EndTime, nil)
	if err != nil {
		log.Fatalf("Failed to insert boat timeblock: %v", err)
	}
	var boatDefaultSettings BoatDefaultSettings

	//read boat DefaultSettings for boatId
	boatDefaultSettings, err = GetDefaultSettingsForBoatId(boatIdString, db)
	if err != nil {
		return 0, err
	}

	//Create boat booking
	query := "INSERT INTO boat_booking (boat_id, booking_id, boat_time_block_id, booking_status_id, booking_file_id, location_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.Exec(query, details.BoatID, details.BookingID, boatTimeblockID, 1, boatDefaultSettings.FileID, details.LocationID)
	if err != nil {
		return 0, err
	}

	//get boat booking id
	boatBookingID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	//Update boat timeblock with boat booking id
	query = "UPDATE boat_timeblock SET boat_booking_id = ? WHERE id = ?"
	_, err = tx.Exec(query, boatBookingID, boatTimeblockID)
	if err != nil {
		return 0, err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		return boatBookingID, err
	}

	return boatBookingID, nil

}

func GetBoatBookingsForBookingId(bookingId string, db *sql.DB) ([]BoatBooking, error) {
	// Query the database for all boat bookings.
	rows, err := db.Query("SELECT * FROM boat_booking WHERE booking_id = ?", bookingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of boat bookings to hold the data.
	var boatBookings []BoatBooking

	// Loop through the data and insert into the boat bookings slice.
	for rows.Next() {
		var boatBooking BoatBooking
		if err := rows.Scan(&boatBooking.ID, &boatBooking.BoatID, &boatBooking.BookingID, &boatBooking.BoatTimeBlockID, &boatBooking.BookingStatusID, &boatBooking.LocationID, &boatBooking.BookingFileID); err != nil {
			return nil, err
		}
		boatBookings = append(boatBookings, boatBooking)
	}

	return boatBookings, nil
}

// API Handlers
func GetBoatBookingsForBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database for all boat bookings.

	boatBookings, err := GetBoatBookingsForBookingId(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatBookings)
}

func GetBoatBookings(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database for all boat bookings.
	rows, err := db.Query("SELECT * FROM boat_booking")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of boat bookings to hold the data.
	var boatBookings []BoatBooking

	// Loop through the data and insert into the boat bookings slice.
	for rows.Next() {
		var boatBooking BoatBooking
		if err := rows.Scan(&boatBooking.ID, &boatBooking.BoatID, &boatBooking.BookingID, &boatBooking.BoatTimeBlockID, &boatBooking.BookingStatusID, &boatBooking.LocationID, &boatBooking.BookingFileID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		boatBookings = append(boatBookings, boatBooking)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatBookings)
}

func CreateBoatBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Decode the request body into a RequestboatBooking struct.
	var details RequestBoatBooking
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	// Attempt to book the boat.
	boatBookingID, err := AttemptToBookBoat(details, db)
	if err != nil {
		log.Fatalf("failed to book boat: %v", err)
	}

	// Return the boat booking ID as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatBookingID)

}

func GetBoatBookingDetails(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database for the boat booking joined with the boat timeblock.
	query := "SELECT rb.id, rb.boat_id, rb.booking_id, rb.boat_time_block_id, rb.booking_status_id, rb.booking_file_id, rt.start_time, rt.end_time, rt.boat_booking_id, rb.location_id FROM boat_booking rb JOIN boat_timeblock rt ON rb.boat_time_block_id = rt.id WHERE rb.id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a single instance of boatBookingDetails.
	var boatBookingDetails BoatBookingDetails

	// Check if there is at least one row.
	if rows.Next() {
		var startTimeStr, endTimeStr string
		var boatBookingID int

		// Scan the values into variables.
		if err := rows.Scan(&boatBookingDetails.ID, &boatBookingDetails.BoatID, &boatBookingDetails.BookingID, &boatBookingDetails.BoatTimeBlockID, &boatBookingDetails.BookingStatusID, &boatBookingDetails.BookingFileID, &startTimeStr, &endTimeStr, &boatBookingID, &boatBookingDetails.LocationID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		// Convert the datetime strings to time.Time.
		boatBookingDetails.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			log.Fatalf("failed to parse start time: %v", err)
		}

		boatBookingDetails.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			log.Fatalf("failed to parse end time: %v", err)
		}

		boatBookingDetails.ID = boatBookingID
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatBookingDetails)

}
