package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type BookingDetails struct {
	ID               int
	BookingID        int
	PaymentComplete  bool
	PaymentDueDate   time.Time
	DocumentsSigned  bool
	BookingStartDate time.Time
}

func GetDetailsForBookingID(bookingId string, db *sql.DB) (BookingDetails, error) {
	rows, err := db.Query("SELECT id, booking_id, payment_complete, payment_due_date, documents_signed, booking_start_date FROM booking_details WHERE booking_id = ?", bookingId)
	if err != nil {
		return BookingDetails{}, err
	}
	defer rows.Close()

	var bookingDetails BookingDetails

	var dueDateString string
	var startDateString string

	if rows.Next() {
		err := rows.Scan(&bookingDetails.ID, &bookingDetails.BookingID, &bookingDetails.PaymentComplete, &dueDateString, &bookingDetails.DocumentsSigned, &startDateString)

		if err != nil {
			return BookingDetails{}, err
		}

		// Attempt to parse with date and time layout
		bookingDetails.PaymentDueDate, err = time.Parse("2006-01-02 15:04:05", dueDateString)
		if err != nil {
			// If parsing fails, attempt to parse with date-only layout
			bookingDetails.PaymentDueDate, err = time.Parse("2006-01-02", dueDateString)
			if err != nil {
				return BookingDetails{}, err
			}
		}

		// Attempt to parse with date and time layout
		bookingDetails.BookingStartDate, err = time.Parse("2006-01-02 15:04:05", startDateString)
		if err != nil {
			// If parsing fails, attempt to parse with date-only layout
			bookingDetails.BookingStartDate, err = time.Parse("2006-01-02", startDateString)
			if err != nil {
				return BookingDetails{}, err
			}
		}

	}

	return bookingDetails, nil
}

func UpdateBookingDetails(bookingDetails BookingDetails, db *sql.DB) error {
	_, err := db.Exec("UPDATE booking_details SET payment_complete = ?, payment_due_date = ?, documents_signed = ?, booking_start_date = ? WHERE id = ?", bookingDetails.PaymentComplete, bookingDetails.PaymentDueDate, bookingDetails.DocumentsSigned, bookingDetails.BookingStartDate, bookingDetails.ID)
	return err
}

func GetBookingDetails(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	bookingDetails, err := GetDetailsForBookingID(id, db)
	if err != nil {
		log.Fatal(err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingDetails)

}

func UpdateBookingDetailsForBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Create a single instance of BookingDetails.
	var bookingDetails BookingDetails

	// Decode the JSON data.
	err := json.NewDecoder(r.Body).Decode(&bookingDetails)
	if err != nil {
		log.Fatal(err)
	}

	// Update the database.
	err = UpdateBookingDetails(bookingDetails, db)
	if err != nil {
		log.Fatal(err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingDetails)

}
