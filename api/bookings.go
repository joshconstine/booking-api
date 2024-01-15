package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Booking struct {
	ID               int
	UserID           int
	BookingStatusID  int
	BookingDetailsID int
}

func createNewBooking(db *sql.DB, userID int) (int, error) {

	//start transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	//create booking
	bookingResult, err := tx.Exec("INSERT INTO booking (user_id, booking_status_id, booking_details_id) VALUES (?, ?, ?)", userID, 1, 0)
	if err != nil {
		log.Fatal(err)
	}

	//get booking id
	bookingID, err := bookingResult.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	tenYearFromNow := time.Now().AddDate(10, 0, 0)

	//create booking details
	bookingResult, err = tx.Exec("INSERT INTO booking_details (booking_id, payment_complete, payment_due_date, documents_signed, booking_start_date) VALUES (?, ?, ?, ?, ?)", bookingID, false, tenYearFromNow, false, tenYearFromNow)
	if err != nil {
		log.Fatal(err)
	}

	//get booking details id
	bookingDetailsID, err := bookingResult.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	//update booking with booking details id
	_, err = tx.Exec("UPDATE booking SET booking_details_id = ? WHERE id = ?", bookingDetailsID, bookingID)
	if err != nil {
		log.Fatal(err)
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return int(bookingID), nil
}

func GetBookings(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM booking")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var bookings []Booking

	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.UserID, &booking.BookingStatusID, &booking.BookingDetailsID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		bookings = append(bookings, booking)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)

}

func CreateBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var booking Booking
	json.NewDecoder(r.Body).Decode(&booking)

	id, err := createNewBooking(db, booking.UserID)

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
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)

}
