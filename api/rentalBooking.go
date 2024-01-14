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

type RentalBooking struct {
	ID                int
	RentalID          int
	BookingID         int
	RentalTimeBlockID int
	BookingStatusID   int
	BookingFileID     int
}

type RequestRentalBooking struct {
	RentalID  int
	BookingID int
	StartTime time.Time
	EndTime   time.Time
}

type RentalBookingDetails struct {
	ID                int
	RentalID          int
	BookingID         int
	RentalTimeBlockID int
	BookingStatusID   int
	BookingFileID     int
	StartTime         time.Time
	EndTime           time.Time
}
type RentalBookingCost struct {
	ID                int
	RentalBookingID   int
	BookingCostItemID int
}

func AddRentalBookingCost(rentalBookingCost RentalBookingCost, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO rental_booking_cost (rental_booking_id, booking_cost_item_id) VALUES (?, ?)", rentalBookingCost.RentalBookingID, rentalBookingCost.BookingCostItemID)
	return err
}

func DetermineCleanFee(rentalUnitDefaultSettings RentalUnitDefaultSettings, rentalUnitVariableSettings []RentalUnitVariableSettings, startTime time.Time, endTime time.Time) float64 {

	// check if cleaning fee is in variable settings
	varSettingFound := false
	for _, rentalUnitVariableSetting := range rentalUnitVariableSettings {
		if startTime.After(rentalUnitVariableSetting.StartDate) && startTime.Before(rentalUnitVariableSetting.EndDate) {
			return rentalUnitVariableSetting.CleaningFee
		}
	}
	// If no variable setting, use default
	if !varSettingFound {
		return rentalUnitDefaultSettings.CleaningFee
	}
	return 0
}

func CalculateRentalCost(rentalUnitVariableSettings []RentalUnitVariableSettings, rentalUnitDefaultSettings RentalUnitDefaultSettings, startTime time.Time, endTime time.Time) float64 {
	// Calculate the duration in days
	durationInDays := int(endTime.Sub(startTime).Hours() / 24)

	// Loop through each day
	var totalCost float64
	for i := 0; i <= durationInDays; i++ {
		currentDay := startTime.Add(time.Duration(i) * 24 * time.Hour)
		// Check if there is a variable setting for this day
		varSettingFound := false
		for _, rentalUnitVariableSetting := range rentalUnitVariableSettings {
			if currentDay.After(rentalUnitVariableSetting.StartDate) && currentDay.Before(rentalUnitVariableSetting.EndDate) {
				totalCost += rentalUnitVariableSetting.NightlyCost
				varSettingFound = true
				break
			}
		}
		// If no variable setting, use default
		if !varSettingFound {
			totalCost += rentalUnitDefaultSettings.NightlyCost
		}
	}

	return totalCost
}

func AttemptToBookRental(details RequestRentalBooking, db *sql.DB) (int64, error) {

	//oprn transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	rentalIdString := strconv.Itoa(details.RentalID)

	//get rental default settings
	rentalSettings, err := GetDefaultSettingsForRentalId(rentalIdString, db)
	if err != nil {
		return 0, err
	}
	// Calculate the duration in days
	durationInDays := details.EndTime.Sub(details.StartTime) / (24 * time.Hour)

	// Check if the duration meets the minimum requirement
	if int(durationInDays) < rentalSettings.MinimumBookingDuration {
		return -2, nil
	}

	// Attempt to create rental timeblock
	rentalTimeblockID, err := AttemptToInsertRentalTimeblock(db, rentalIdString, details.StartTime, details.EndTime, nil)
	if err != nil {
		log.Fatalf("Failed to insert rental timeblock: %v", err)
	}

	if rentalTimeblockID == -1 {
		tx.Rollback()
		return -1, nil
	}

	var rentalUnitDefaultSettings RentalUnitDefaultSettings

	//read rental DefaultSettings for rentalId
	rentalUnitDefaultSettings, err = GetDefaultSettingsForRentalId(rentalIdString, db)
	if err != nil {
		return 0, err
	}

	//Create rental booking
	query := "INSERT INTO rental_booking (rental_id, booking_id, rental_time_block_id, booking_status_id, booking_file_id) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.Exec(query, details.RentalID, details.BookingID, rentalTimeblockID, 1, rentalUnitDefaultSettings.FileID)
	if err != nil {
		return 0, err
	}

	//get rental booking id
	rentalBookingID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	//Update rental timeblock with rental booking id
	query = "UPDATE rental_timeblock SET rental_booking_id = ? WHERE id = ?"
	_, err = tx.Exec(query, rentalBookingID, rentalTimeblockID)
	if err != nil {
		return 0, err
	}

	//get variable settings for Dates
	var rentalUnitVariableSettings []RentalUnitVariableSettings

	rentalUnitVariableSettings, err = GetVariableSettingsForRentalIdAndDateRange(rentalIdString, details.StartTime, details.EndTime, db)
	if err != nil {
		return 0, err
	}

	//calculate cost
	totalCost := CalculateRentalCost(rentalUnitVariableSettings, rentalUnitDefaultSettings, details.StartTime, details.EndTime)

	//create booking cost item for rental cost
	rentalFeeBookingCostItem := BookingCostItem{
		BookingID:         details.BookingID,
		BookingCostTypeID: 3,
		Ammount:           totalCost,
	}
	//create Cleanign fee cost item

	cleaningFee := DetermineCleanFee(rentalUnitDefaultSettings, rentalUnitVariableSettings, details.StartTime, details.EndTime)

	cleaningFeeBookingCostItem := BookingCostItem{
		BookingID:         details.BookingID,
		BookingCostTypeID: 2,
		Ammount:           cleaningFee,
	}

	createdRentalFee, err := AttemptToCreateBookingCostItem(rentalFeeBookingCostItem, db)
	if err != nil {
		return 0, err
	}

	createdCleaningFee, err := AttemptToCreateBookingCostItem(cleaningFeeBookingCostItem, db)
	if err != nil {
		return 0, err
	}

	//create rental booking cost
	rentalBookingCost := RentalBookingCost{
		RentalBookingID:   int(rentalBookingID),
		BookingCostItemID: createdRentalFee,
	}

	err = AddRentalBookingCost(rentalBookingCost, db)
	if err != nil {
		return 0, err
	}

	//create cleaning fee booking cost
	cleaningFeeBookingCost := RentalBookingCost{
		RentalBookingID:   int(rentalBookingID),
		BookingCostItemID: createdCleaningFee,
	}

	err = AddRentalBookingCost(cleaningFeeBookingCost, db)
	if err != nil {
		return 0, err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		return rentalBookingID, err
	}

	return rentalBookingID, nil

}

func GetRentalBookingsForBookingId(bookingId string, db *sql.DB) ([]RentalBooking, error) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT * FROM rental_booking WHERE booking_id = ?", bookingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var rentalBookings []RentalBooking

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var rentalBooking RentalBooking
		if err := rows.Scan(&rentalBooking.ID, &rentalBooking.RentalID, &rentalBooking.BookingID, &rentalBooking.RentalTimeBlockID, &rentalBooking.BookingStatusID, &rentalBooking.BookingFileID); err != nil {
			return nil, err
		}
		rentalBookings = append(rentalBookings, rentalBooking)
	}

	return rentalBookings, nil
}

// API Handlers
func GetRentalBookingsForBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database for all rental bookings.

	rentalBookings, err := GetRentalBookingsForBookingId(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalBookings)
}

func GetRentalBookings(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT * FROM rental_booking")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var rentalBookings []RentalBooking

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var rentalBooking RentalBooking
		if err := rows.Scan(&rentalBooking.ID, &rentalBooking.RentalID, &rentalBooking.BookingID, &rentalBooking.RentalTimeBlockID, &rentalBooking.BookingStatusID, &rentalBooking.BookingFileID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		rentalBookings = append(rentalBookings, rentalBooking)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalBookings)
}

func CreateRentalBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Decode the request body into a RequestRentalBooking struct.
	var details RequestRentalBooking
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	// Attempt to book the rental.
	rentalBookingID, err := AttemptToBookRental(details, db)
	if err != nil {
		log.Fatalf("failed to book rental: %v", err)
	}

	if rentalBookingID == -1 {
		// Return a 409 Conflict if the rental is already booked.
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Rental is already booked"))

		return
	}

	if rentalBookingID == -2 {
		//minimum duration not met
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Minimum duration not met"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	// Return the rental booking ID as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalBookingID)

}

func GetRentalBookingDetails(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	rentalBookingId := vars["rentalBookingId"]

	// Query the database for the rental booking joined with the rental timeblock.
	query := "SELECT rb.id, rb.rental_id, rb.booking_id, rb.rental_time_block_id, rb.booking_status_id, rb.booking_file_id, rt.start_time, rt.end_time, rt.rental_booking_id FROM rental_booking rb JOIN rental_timeblock rt ON rb.rental_time_block_id = rt.id WHERE rb.id = ?"
	rows, err := db.Query(query, rentalBookingId)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a single instance of RentalBookingDetails.
	var rentalBookingDetails RentalBookingDetails

	// Check if there is at least one row.
	if rows.Next() {
		var startTimeStr, endTimeStr string
		var rentalBookingID int
		// Scan the values into variables.
		if err := rows.Scan(&rentalBookingDetails.ID, &rentalBookingDetails.RentalID, &rentalBookingDetails.BookingID, &rentalBookingDetails.RentalTimeBlockID, &rentalBookingDetails.BookingStatusID, &rentalBookingDetails.BookingFileID, &startTimeStr, &endTimeStr, &rentalBookingID); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		// Convert the datetime strings to time.Time.
		rentalBookingDetails.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			log.Fatalf("failed to parse start time: %v", err)
		}

		rentalBookingDetails.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			log.Fatalf("failed to parse end time: %v", err)
		}

		rentalBookingDetails.ID = rentalBookingID
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalBookingDetails)

}
