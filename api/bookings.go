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

type Booking struct {
	ID               int
	UserID           int
	BookingStatusID  int
	BookingDetailsID int
}

type BookingInformation struct {
	BookingID      int
	BookingStatus  BookingStatus
	User           User
	BookingDetails BookingDetails
	RentalBookings []RentalBookingDetails
	CostItems      []BookingCostItem
	Payments       []BookingPayment
}

type BookingSnapshot struct {
	BookingID       int
	BookingStatus   string
	RentalsBooked   []string
	BoatsBooked     []string
	Events          []string
	HasAlcoholOrder bool
	BookingDetails  BookingDetails
	User            User
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
func GetInformationForBookingID(bookingId string, db *sql.DB) (BookingInformation, error) {

	//get booking
	rows, err := db.Query("SELECT b.id, bs.name, bd.id, bd.payment_complete, bd.payment_due_date, bd.documents_signed, bd.booking_start_date, u.id, u.first_name, u.last_name, u.email, u.phone_number, b.booking_status_id FROM booking b JOIN booking_status bs ON b.booking_status_id = bs.id JOIN booking_details bd ON b.booking_details_id = bd.id JOIN user u ON b.user_id = u.id WHERE b.id = ?", bookingId)
	if err != nil {
		return BookingInformation{}, err
	}
	defer rows.Close()

	var bookingInformation BookingInformation

	var dueDateString string
	var startDateString string

	if rows.Next() {
		err := rows.Scan(&bookingInformation.BookingID, &bookingInformation.BookingStatus.Name, &bookingInformation.BookingDetails.ID, &bookingInformation.BookingDetails.PaymentComplete, &dueDateString, &bookingInformation.BookingDetails.DocumentsSigned, &startDateString, &bookingInformation.User.ID, &bookingInformation.User.FirstName, &bookingInformation.User.LastName, &bookingInformation.User.Email, &bookingInformation.User.PhoneNumber, &bookingInformation.BookingStatus.ID)

		if err != nil {
			return BookingInformation{}, err
		}

		// Attempt to parse with date and time layout
		bookingInformation.BookingDetails.PaymentDueDate, err = time.Parse("2006-01-02 15:04:05", dueDateString)
		if err != nil {
			// If parsing fails, attempt to parse with date-only layout
			bookingInformation.BookingDetails.PaymentDueDate, err = time.Parse("2006-01-02", dueDateString)
			if err != nil {
				return BookingInformation{}, err
			}
		}

		// Attempt to parse with date and time layout
		bookingInformation.BookingDetails.BookingStartDate, err = time.Parse("2006-01-02 15:04:05", startDateString)
		if err != nil {
			// If parsing fails, attempt to parse with date-only layout
			bookingInformation.BookingDetails.BookingStartDate, err = time.Parse("2006-01-02", startDateString)
			if err != nil {
				return BookingInformation{}, err
			}
		}

	}

	//get rental bookings
	var rentalBookings []RentalBookingDetails

	rentalBookingIds, err := GetRentalBookingIDsForBookingId(bookingId, db)

	for _, rentalBookingId := range rentalBookingIds {
		rbIdString := strconv.Itoa(rentalBookingId)
		rentalBooking, err := GetDetailsForRentalBookingID(rbIdString, db)
		if err != nil {
			return BookingInformation{}, err
		}
		rentalBookings = append(rentalBookings, rentalBooking)
	}

	bookingInformation.RentalBookings = rentalBookings

	//get cost items
	costItems, err := GetCostItemsForBookingId(bookingId, db)
	if err != nil {
		return BookingInformation{}, err
	}

	bookingInformation.CostItems = costItems

	bookingIdInt, err := strconv.Atoi(bookingId)
	if err != nil {
		return BookingInformation{}, err
	}
	//get payments
	payments, err := GetBookingPaymentsForBookingID(bookingIdInt, db)
	if err != nil {
		return BookingInformation{}, err
	}

	bookingInformation.Payments = payments
	return bookingInformation, nil

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
func GetBookingInformation(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	id := vars["id"]

	bookingInformation, err := GetInformationForBookingID(id, db)
	if err != nil {
		log.Fatal(err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingInformation)

}

func GetBookingSnapshots(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	//get query params
	vars := r.URL.Query()
	startDateString := vars.Get("startDate")
	endDateString := vars.Get("endDate")

	//parse dates or set to defaults today, 10 years from now
	var startDate time.Time
	var endDate time.Time

	if startDateString == "" {
		startDate = time.Now()
	} else {
		startDate, _ = time.Parse("2006-01-02", startDateString)
	}

	if endDateString == "" {
		endDate = time.Now().AddDate(10, 0, 0)
	} else {
		endDate, _ = time.Parse("2006-01-02", endDateString)
	}

	// perform query

	rows, err := db.Query("SELECT b.id,b.user_id, bs.name, bd.id, bd.payment_complete, bd.payment_due_date, bd.documents_signed, bd.booking_start_date FROM booking b JOIN booking_status bs ON b.booking_status_id = bs.id JOIN booking_details bd ON b.booking_details_id = bd.id WHERE bd.booking_start_date >= ? AND bd.booking_start_date <= ?", startDate, endDate)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var bookingSnapshots []BookingSnapshot

	for rows.Next() {
		var user User

		var bookingSnapshot BookingSnapshot

		var dueDateString string
		var startDateString string

		if err := rows.Scan(&bookingSnapshot.BookingID, &bookingSnapshot.User.ID, &bookingSnapshot.BookingStatus, &bookingSnapshot.BookingDetails.ID, &bookingSnapshot.BookingDetails.PaymentComplete, &endDateString, &bookingSnapshot.BookingDetails.DocumentsSigned, &startDateString); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		// Attempt to parse with date and time layout
		bookingSnapshot.BookingDetails.PaymentDueDate, err = time.Parse("2006-01-02 15:04:05", dueDateString)
		if err != nil {
		}

		// Attempt to parse with date and time layout
		bookingSnapshot.BookingDetails.BookingStartDate, err = time.Parse("2006-01-02 15:04:05", startDateString)
		if err != nil {

		}

		//get user
		user, err := GetUserForUserID(strconv.Itoa(bookingSnapshot.User.ID), db)
		if err != nil {
			log.Fatal(err)
		}

		rentalNames, err := GetRentalNamesForBookingId(strconv.Itoa(bookingSnapshot.BookingID), db)
		if err != nil {
			log.Fatal(err)
		}

		bookingSnapshot.User = user
		bookingSnapshot.RentalsBooked = rentalNames

		bookingSnapshots = append(bookingSnapshots, bookingSnapshot)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookingSnapshots)
}
