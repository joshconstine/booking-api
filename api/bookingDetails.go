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

func GetBookingDetails(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database.
	rows, err := db.Query("SELECT * FROM booking_details WHERE booking_id = ?", id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var bookingDetails BookingDetails

	var dueDateString string
	var startDateString string

	if rows.Next() {
		err := rows.Scan(&bookingDetails.ID, &bookingDetails.BookingID, &bookingDetails.PaymentComplete, &dueDateString, &bookingDetails.DocumentsSigned, &startDateString)

		if err != nil {
			log.Fatalf("failed to scan: %v", err)
		}

		// Attempt to parse with date and time layout
		bookingDetails.PaymentDueDate, err = time.Parse("2006-01-02 15:04:05", dueDateString)
		if err != nil {
			// If parsing fails, attempt to parse with date-only layout
			bookingDetails.PaymentDueDate, err = time.Parse("2006-01-02", dueDateString)
			if err != nil {
				log.Fatalf("failed to parse date: %v", err)
			}
		}

		// Attempt to parse with date and time layout
		bookingDetails.BookingStartDate, err = time.Parse("2006-01-02 15:04:05", startDateString)
		if err != nil {
			// If parsing fails, attempt to parse with date-only layout
			bookingDetails.BookingStartDate, err = time.Parse("2006-01-02", startDateString)
			if err != nil {
				log.Fatalf("failed to parse date: %v", err)
			}
		}

	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingDetails)
}
