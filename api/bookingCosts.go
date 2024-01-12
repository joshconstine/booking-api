package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type BookingCostType struct {
	ID   int
	Name string
}

func GetBookingCostTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM booking_cost_type")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var bookingCostTypes []BookingCostType

	for rows.Next() {
		var bookingCostType BookingCostType
		if err := rows.Scan(&bookingCostType.ID, &bookingCostType.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		bookingCostTypes = append(bookingCostTypes, bookingCostType)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingCostTypes)
}
func CreateBookingCostType(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var bookingCostType BookingCostType
	if err := json.NewDecoder(r.Body).Decode(&bookingCostType); err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	_, err := db.Exec("INSERT INTO booking_cost_type (name) VALUES (?)", bookingCostType.Name)
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
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Function to check if the error is a duplicate key error
func IsDuplicateKeyError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		return true
	}
	return false
}
