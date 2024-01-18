package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BookingPayment struct {
	ID              int
	BookingID       int
	PaymentAmount   float64
	PaymentMethodID int
	PaypalOrderID   *string
}

func GetAllBookingPayments(db *sql.DB) ([]BookingPayment, error) {

	var bookingPayments []BookingPayment

	rows, err := db.Query("SELECT id, booking_id, payment_amount, payment_method_id, paypal_order_id FROM booking_payment")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var bookingPayment BookingPayment

		err := rows.Scan(&bookingPayment.ID, &bookingPayment.BookingID, &bookingPayment.PaymentAmount, &bookingPayment.PaymentMethodID, &bookingPayment.PaypalOrderID)

		if err != nil {
			log.Fatalf("failed to query: %v", err)
		}

		bookingPayments = append(bookingPayments, bookingPayment)
	}

	return bookingPayments, err
}

func GetBookingPaymentsForBookingID(bookingID int, db *sql.DB) ([]BookingPayment, error) {

	var bookingPayments []BookingPayment

	rows, err := db.Query("SELECT id, booking_id, payment_amount, payment_method_id, paypal_order_id FROM booking_payment WHERE booking_id = ?", bookingID)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var bookingPayment BookingPayment

		err := rows.Scan(&bookingPayment.ID, &bookingPayment.BookingID, &bookingPayment.PaymentAmount, &bookingPayment.PaymentMethodID, &bookingPayment.PaypalOrderID)

		if err != nil {
			log.Fatalf("failed to query: %v", err)
		}

		bookingPayments = append(bookingPayments, bookingPayment)
	}

	return bookingPayments, err
}
func GetTotalPaymentsForBookingID(bookingID int, db *sql.DB) (float64, error) {

	var totalPayments float64

	rows, err := db.Query("SELECT SUM(payment_amount) FROM booking_payment WHERE booking_id = ?", bookingID)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&totalPayments)

		if err != nil {
			log.Fatalf("failed to query: %v", err)
		}

	}

	return totalPayments, err
}

func AddPaymentToBooking(bookingPayment BookingPayment, db *sql.DB) (int64, error) {

	result, err := db.Exec("INSERT INTO booking_payment (booking_id, payment_amount, payment_method_id, paypal_order_id) VALUES (?, ?, ?, ?)", bookingPayment.BookingID, bookingPayment.PaymentAmount, bookingPayment.PaymentMethodID, bookingPayment.PaypalOrderID)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	//update the booking details
	_, err = VerifyBookingPaymentStatus(bookingPayment.BookingID, db)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	return result.LastInsertId()
}

func GetBookingPayments(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	bookingPayments, err := GetAllBookingPayments(db)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingPayments)

}

// Handling JSON decoding error, sending appropriate HTTP response
func CreateBookingPayment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var bookingPayment BookingPayment
	if err := json.NewDecoder(r.Body).Decode(&bookingPayment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := AddPaymentToBooking(bookingPayment, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to add payment: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
