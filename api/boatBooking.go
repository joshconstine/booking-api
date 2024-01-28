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

	//get boat default settings

	boatSettings, err := GetDefaultSettingsForBoatId(boatIdString, db)
	if err != nil {
		return 0, err
	}

	boatStatus, err := GetStatusForBoatId(boatIdString, db)
	if err != nil {
		return 0, err
	}

	if boatStatus.CurrentLocationID != details.LocationID {
		if boatSettings.AdvertiseAtAllLocations == false {
			tx.Rollback()
			return -2, nil
		}
	}

	// Attempt to create boat timeblock
	boatTimeblockID, err := AttemptToInsertBoatTimeblock(db, details.BoatID, details.StartTime, details.EndTime, nil)
	if err != nil {
		log.Fatalf("Failed to insert boat timeblock: %v", err)
	}

	if boatTimeblockID == -1 {
		tx.Rollback()
		return -1, nil
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

	if boatBookingID == -1 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Boat is already booked for this time"))
		return
	}

	if boatBookingID == -2 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Boat is not available at this location"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	// Return the boat booking ID as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatBookingID)

}

func GetBoatBookingDetails(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	boatBookingId := vars["boatBookingId"]

	query := "SELECT boat_booking.id, boat_booking.boat_id, boat_booking.booking_id, boat_booking.boat_time_block_id, boat_booking.booking_status_id, boat_booking.booking_file_id, boat_booking.location_id, boat_timeblock.start_time, boat_timeblock.end_time FROM boat_booking JOIN boat_timeblock ON boat_booking.boat_time_block_id = boat_timeblock.id WHERE boat_booking.id = ?"

	// Query the database.
	rows, err := db.Query(query, boatBookingId)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	var boatBookingDetails BoatBookingDetails

	if rows.Next() {
		var startTimeStr, endTimeStr string

		// Scan the values into variables.
		if err := rows.Scan(&boatBookingDetails.ID, &boatBookingDetails.BoatID, &boatBookingDetails.BookingID, &boatBookingDetails.BoatTimeBlockID, &boatBookingDetails.BookingStatusID, &boatBookingDetails.BookingFileID, &boatBookingDetails.LocationID, &startTimeStr, &endTimeStr); err != nil {
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
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatBookingDetails)
}
