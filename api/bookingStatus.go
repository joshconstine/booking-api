package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type BookingStatus struct {
	ID   int
	Name string
}

func GetBookingStatus(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM booking_status")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var bookingStatus []BookingStatus

	for rows.Next() {
		var status BookingStatus
		if err := rows.Scan(&status.ID, &status.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		bookingStatus = append(bookingStatus, status)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingStatus)

}
