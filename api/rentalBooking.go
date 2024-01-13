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

type RentalBooking struct {
	ID                int
	RentalID          int
	BookingID         int
	RentalTimeBlockID int
	BookingStatusID   int
	BookingFileID     int
}

type RequestRentalBooking struct {
	RentalID  int
	BookingID int
	StartTime time.Time
	EndTime   time.Time
}

func AttemptToBookRental(details RequestRentalBooking, db *sql.DB) (int64, error) {

	//oprn transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	rentalIdString := strconv.Itoa(details.RentalID)

	// Attempt to create rental timeblock
	rentalTimeblockID, err := AttemptToInsertRentalTimeblock(db, rentalIdString, details.StartTime, details.EndTime, nil)
	if err != nil {
		log.Fatalf("Failed to insert rental timeblock: %v", err)
	}
	var rentalUnitDefaultSettings RentalUnitDefaultSettings

	//read rental DefaultSettings for rentalId
	rentalUnitDefaultSettings, err = GetDefaultSettingsForRentalId(rentalIdString, db)
	if err != nil {
		return 0, err
	}

	//Create rental booking
	query := "INSERT INTO rental_booking (rental_id, booking_id, rental_time_block_id, booking_status_id, booking_file_id) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.Exec(query, details.RentalID, details.BookingID, rentalTimeblockID, 1, rentalUnitDefaultSettings.FileID)
	if err != nil {
		return 0, err
	}

	//get rental booking id
	rentalBookingID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	//Update rental timeblock with rental booking id
	query = "UPDATE rental_timeblock SET rental_booking_id = ? WHERE id = ?"
	_, err = tx.Exec(query, rentalBookingID, rentalTimeblockID)
	if err != nil {
		return 0, err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		return rentalBookingID, err
	}

	return rentalBookingID, nil

}

func GetRentalBookingsForBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT * FROM rental_booking WHERE booking_id = ?", id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var rentalBookings []RentalBooking

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var rentalBooking RentalBooking
		if err := rows.Scan(&rentalBooking.ID, &rentalBooking.RentalID, &rentalBooking.BookingID, &rentalBooking.RentalTimeBlockID, &rentalBooking.BookingStatusID, &rentalBooking.BookingFileID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		rentalBookings = append(rentalBookings, rentalBooking)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalBookings)
}

func GetRentalBookings(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT * FROM rental_booking")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var rentalBookings []RentalBooking

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var rentalBooking RentalBooking
		if err := rows.Scan(&rentalBooking.ID, &rentalBooking.RentalID, &rentalBooking.BookingID, &rentalBooking.RentalTimeBlockID, &rentalBooking.BookingStatusID, &rentalBooking.BookingFileID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		rentalBookings = append(rentalBookings, rentalBooking)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalBookings)
}

func CreateRentalBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Decode the request body into a RequestRentalBooking struct.
	var details RequestRentalBooking
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	// Attempt to book the rental.
	rentalBookingID, err := AttemptToBookRental(details, db)
	if err != nil {
		log.Fatalf("failed to book rental: %v", err)
	}

	// Return the rental booking ID as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalBookingID)

}
