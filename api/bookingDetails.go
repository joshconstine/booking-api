package api

import (
	"booking-api/api/payments"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	InvoiceID        *string
}

func VerifyBookingPaymentStatus(bookingID int, db *sql.DB) (bool, error) {

	bookingIdString := strconv.Itoa(bookingID)

	bookingDetails, err := GetDetailsForBookingID(bookingIdString, db)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	totalBill, err := GetTotalCostItemsForBookingID(bookingID, db)
	if err != nil {
		return false, fmt.Errorf("failed to query total bill: %v", err)
	}

	totalPayments, err := GetTotalPaymentsForBookingID(bookingID, db)
	if err != nil {
		return false, fmt.Errorf("failed to query total payments: %v", err)
	}

	paymentComplete := totalPayments >= totalBill

	// Update only if there is a change in the payment status
	if bookingDetails.PaymentComplete != paymentComplete {
		bookingDetails.PaymentComplete = paymentComplete
		err = UpdateBookingDetails(bookingDetails, db)
		if err != nil {
			return false, fmt.Errorf("failed to update booking details: %v", err)
		}
	}

	return paymentComplete, err
}

func GetDetailsForBookingID(bookingId string, db *sql.DB) (BookingDetails, error) {
	rows, err := db.Query("SELECT id, booking_id, payment_complete, payment_due_date, documents_signed, booking_start_date, invoice_id FROM booking_details WHERE booking_id = ?", bookingId)
	if err != nil {
		return BookingDetails{}, err
	}
	defer rows.Close()

	var bookingDetails BookingDetails

	var dueDateString string
	var startDateString string

	if rows.Next() {
		err := rows.Scan(&bookingDetails.ID, &bookingDetails.BookingID, &bookingDetails.PaymentComplete, &dueDateString, &bookingDetails.DocumentsSigned, &startDateString, &bookingDetails.InvoiceID)

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
		// Log the error and send a bad request response
		log.Printf("error decoding request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
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

func HandleGetInvoiceForBookingId(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	bookingDetails, err := GetDetailsForBookingID(id, db)
	if err != nil {
		log.Fatal(err)
	}

	if bookingDetails.InvoiceID == nil {
		// http.Error(w, "No invoice found for booking", http.StatusNotFound)
		// return
		// bookingDetails.InvoiceID = "INV2-YD5J-39TL-ZPKD-J8JR"
		bookingDetails.InvoiceID = new(string)
		*bookingDetails.InvoiceID = "INV2-YD5J-39TL-ZPKD-J8JR"

	}

	client := payments.CreatePaypalClient()
	invoice, err := payments.GetInvoiceByID(r.Context(), client, *bookingDetails.InvoiceID)
	if err != nil {
		http.Error(w, "Failed to get invoice details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)

}
