package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BookingCostItem struct {
	ID                int
	BookingID         int
	BookingCostTypeID int
	Ammount           float64
}

func GetCostItemsForBookingId(bookingId string, db *sql.DB) ([]BookingCostItem, error) {
	rows, err := db.Query("SELECT id, booking_id, booking_cost_type_id, ammount FROM booking_cost_item WHERE booking_id = ?", bookingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookingCostItems []BookingCostItem

	for rows.Next() {
		var bookingCostItem BookingCostItem

		err := rows.Scan(
			&bookingCostItem.ID,
			&bookingCostItem.BookingID,
			&bookingCostItem.BookingCostTypeID,
			&bookingCostItem.Ammount,
		)
		if err != nil {
			return nil, err
		}

		bookingCostItems = append(bookingCostItems, bookingCostItem)
	}

	return bookingCostItems, nil
}

func AttemptToCreateBookingCostItem(bookingCostItem BookingCostItem, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO booking_cost_item (booking_id, booking_cost_type_id, ammount) VALUES (?, ?, ?)", bookingCostItem.BookingID, bookingCostItem.BookingCostTypeID, bookingCostItem.Ammount)
	return err
}

func GetBookingCostItems(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	bookingId := vars["id"]

	bookingCostItems, err := GetCostItemsForBookingId(bookingId, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingCostItems)
}

func CreateBookingCostItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var bookingCostItem BookingCostItem
	if err := json.NewDecoder(r.Body).Decode(&bookingCostItem); err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	err := AttemptToCreateBookingCostItem(bookingCostItem, db)
	if err != nil {
		// Check if the error is a duplicate entry error
		if IsDuplicateKeyError(err) {
			// Handle duplicate entry error
			w.WriteHeader(http.StatusConflict) // HTTP 409 Conflict
			w.Write([]byte("Duplicate entry: The booking cost item already exists."))
		} else {
			// Handle other errors
			log.Printf("failed to insert: %v", err)
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500 Internal Server Error
			w.Write([]byte("Internal Server Error"))
		}
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	w.Write([]byte("Created"))
}
