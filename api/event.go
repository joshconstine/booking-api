package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Event struct {
	ID               int
	Name             string
	BookingID        int
	VenueEventTypeID int
	VenueTimeBlockID int
}

type RequestEvent struct {
	VenueEventTypeID int
	VenueID          int
	BookingID        int
	StartTime        time.Time
	EndTime          time.Time
}

type EventBookingCost struct {
	ID                int
	EventID           int
	BookingCostItemID int
}
type EventDetails struct {
	ID               int
	VenueID          int
	BookingID        int
	VenueTimeBlockID int
	StartTime        time.Time
	EndTime          time.Time
	CostItems        []BookingCostItem
}

func GetNameFromVenueEventTypeID(venueEventTypeID int, db *sql.DB) string {
	var name string
	err := db.QueryRow("SELECT et.name FROM venue_event_type vet  JOIN event_type et ON vet.event_type_id = et.id WHERE vet.id = ?", venueEventTypeID).Scan(&name)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	return name
}

func GetDetailsForEventID(eventId string, db *sql.DB) (EventDetails, error) {

	// Query the database for the event  joined with the event timeblock.

	query := "SELECT e.id, vet.venue_id, e.booking_id, vtb.start_time, vtb.end_time FROM event e JOIN venue_timeblock vtb ON e.venue_timeblock_id = vtb.id JOIN venue_event_type vet ON e.venue_event_type_id = vet.id WHERE e.id = ?"
	rows, err := db.Query(query, eventId)
	if err != nil {
		return EventDetails{}, err
	}
	defer rows.Close()

	// Create a single instance of EventDetails.
	var eventDetails EventDetails

	// Check if there is at least one row.
	if rows.Next() {
		var startTimeStr, endTimeStr string
		var eventID int
		var costItems []BookingCostItem

		// Scan the row into the EventDetails struct.
		if err := rows.Scan(&eventID, &eventDetails.VenueID, &eventDetails.BookingID, &eventDetails.VenueTimeBlockID, &startTimeStr, &endTimeStr); err != nil {
			return EventDetails{}, err
		}

		// Parse the start and end times into time.Time.
		eventDetails.StartTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return EventDetails{}, err
		}
		eventDetails.EndTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			return EventDetails{}, err
		}

		// Get the cost items for the event.
		costItems, err = GetCostItemsForEventId(eventId, db)
		if err != nil {
			return EventDetails{}, err
		}

		eventDetails.CostItems = costItems

	} else {
		return EventDetails{}, nil

	}

	return eventDetails, nil

}

func GetEventIdsFromVenueIDWithRange(venueID int, from time.Time, to time.Time, db *sql.DB) ([]int, error) {

	rentalIdString := strconv.Itoa(venueID)
	fromString := from.Format("2006-01-02 15:04:05")
	toString := to.Format("2006-01-02 15:04:05")

	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT rb.id FROM event rb JOIN rental_timeblock rt ON rb.rental_time_block_id = rt.id WHERE rb.rental_id = ? AND rt.start_time >= ? AND rt.end_time <= ?", rentalIdString, fromString, toString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookingIds []int

	for rows.Next() {
		var bookingId int
		if err := rows.Scan(&bookingId); err != nil {
			return nil, err
		}
		bookingIds = append(bookingIds, bookingId)
	}

	return bookingIds, nil

}

func GetEventDetailsByVenueIdForRange(venueID int, from time.Time, to time.Time, db *sql.DB) ([]EventDetails, error) {

	eventIds, err := GetEventIdsFromVenueIDWithRange(venueID, from, to, db)
	if err != nil {
		return nil, err
	}

	var eventDetails []EventDetails

	for _, eventId := range eventIds {
		eventDetail, err := GetDetailsForEventID(strconv.Itoa(eventId), db)
		if err != nil {
			return nil, err
		}
		eventDetails = append(eventDetails, eventDetail)
	}

	return eventDetails, nil

}

func AddEventBookingCost(eventCost EventBookingCost, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO event_booking_cost (event_id, booking_cost_item_id) VALUES (?, ?)", eventCost.EventID, eventCost.BookingCostItemID)
	return err
}
func GetCostItemsForEventId(eventId string, db *sql.DB) ([]BookingCostItem, error) {

	rows, err := db.Query("SELECT bci.id, bci.booking_id, bci.booking_cost_type_id, bci.amount FROM booking_cost_item bci JOIN event_booking_cost rbc ON bci.id = rbc.booking_cost_item_id WHERE rbc.event_id = ?", eventId)
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
			&bookingCostItem.Amount,
		)
		if err != nil {
			return nil, err
		}

		bookingCostItems = append(bookingCostItems, bookingCostItem)
	}

	return bookingCostItems, nil
}

func CalculateEventHourlyCost(venueEventTypeDefaultSettings VenueEventTypeDefaultSettings, startTime time.Time, endTime time.Time) float64 {
	// Calculate the duration in days
	durationInHours := int(endTime.Sub(startTime).Hours())

	// Calculate the cost based on the hourly rate and flat fee.

	cost := float64(durationInHours) * float64(venueEventTypeDefaultSettings.HourlyRate)

	if cost < float64(venueEventTypeDefaultSettings.FlatFee) {
		cost = float64(venueEventTypeDefaultSettings.FlatFee)
	}

	return cost

}

func AttemptToCreateEvent(details RequestEvent, db *sql.DB) (int64, error) {

	//oprn transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	venueEventTypeIDString := strconv.Itoa(details.VenueEventTypeID)
	//get rental default settings
	venueEventTypeDefaultSettings, err := GetDefaultSettingsForVenueEventTypeID(venueEventTypeIDString, db)
	if err != nil {
		return 0, err
	}

	durationInHours := details.EndTime.Sub(details.StartTime).Hours()
	// Check if the duration meets the minimum requirement
	if int(durationInHours) < venueEventTypeDefaultSettings.MinimumBookingDuration {
		return -2, nil
	}

	// Attempt to create rental timeblock
	// rentalTimeblockID, err := AttemptToInsertRentalTimeblock(db, rentalIdString, details.StartTime, details.EndTime, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to insert rental timeblock: %v", err)
	// }

	venueTimeblockID, err := AttemptToInsertVenueTimeblock(db, details.VenueID, details.StartTime, details.EndTime, nil, nil)
	if err != nil {
		log.Fatalf("Failed to insert venue timeblock: %v", err)
	}

	if venueTimeblockID == -1 {
		tx.Rollback()
		return -1, nil
	}

	name := GetNameFromVenueEventTypeID(details.VenueEventTypeID, db)

	//Create event
	query := "INSERT INTO event (name, venue_event_type_id, booking_id, venue_timeblock_id) VALUES (?, ?, ?, ?)"
	result, err := tx.Exec(query, name, details.VenueEventTypeID, details.BookingID, venueTimeblockID)
	if err != nil {
		return 0, err
	}

	//get event id
	eventID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	//Update venue timeblock with rental booking id
	query = "UPDATE venue_timeblock SET event_id = ? WHERE id = ?"
	_, err = tx.Exec(query, eventID, venueTimeblockID)
	if err != nil {
		return 0, err
	}

	//get booking details
	bookingDetails, err := GetDetailsForBookingID(strconv.Itoa(details.BookingID), db)
	if err != nil {
		return 0, err
	}

	//update booking details
	bookingDetails.BookingStartDate = details.StartTime

	twoWeeksBeforeBookingStartDate := details.StartTime.AddDate(0, 0, -14)

	bookingDetails.PaymentDueDate = twoWeeksBeforeBookingStartDate

	err = UpdateBookingDetails(bookingDetails, db)

	if err != nil {
		return 0, err
	}

	//calculate cost
	totalCost := CalculateEventHourlyCost(venueEventTypeDefaultSettings, details.StartTime, details.EndTime)

	//create booking cost item for rental cost
	eventFeeBookingCostItem := BookingCostItem{
		BookingID:         details.BookingID,
		BookingCostTypeID: 9,
		Amount:            totalCost,
	}
	createdEventFee, err := AttemptToCreateBookingCostItem(eventFeeBookingCostItem, db)
	if err != nil {
		return 0, err
	}

	//create rental booking cost
	eventCost := EventBookingCost{
		EventID:           int(eventID),
		BookingCostItemID: createdEventFee,
	}

	err = AddEventBookingCost(eventCost, db)
	if err != nil {
		return 0, err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		return eventID, err
	}

	return eventID, nil

}

func GetEventsForBookingId(bookingId string, db *sql.DB) ([]Event, error) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT id, name, booking_id, venue_event_type_id, venue_timeblock_id FROM event WHERE booking_id = ?", bookingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var events []Event

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.BookingID, &event.VenueEventTypeID, &event.VenueTimeBlockID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventsForVenueId(rentalId int, db *sql.DB) ([]Event, error) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT id, name, booking_id, venue_event_type_id, rental_time_block_id FROM event WHERE rental_id = ?", rentalId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var events []Event

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.BookingID, &event.VenueEventTypeID, &event.VenueTimeBlockID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
func GetEventIDsForBookingId(bookingId string, db *sql.DB) ([]int, error) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT id FROM event WHERE booking_id = ?", bookingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var eventIds []int

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var eventId int
		if err := rows.Scan(&eventId); err != nil {
			return nil, err
		}
		eventIds = append(eventIds, eventId)
	}

	return eventIds, nil
}

func GetEventNamesForBookingId(bookingId string, db *sql.DB) ([]string, error) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT name FROM event WHERE booking_id = ?", bookingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var eventNames []string

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var eventName string
		if err := rows.Scan(&eventName); err != nil {
			return nil, err
		}
		eventNames = append(eventNames, eventName)
	}

	return eventNames, nil
}

// API Handlers
func GetEventsForBooking(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database for all rental bookings.

	events, err := GetEventsForBookingId(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func GetEvents(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database for all rental bookings.
	rows, err := db.Query("SELECT id, name, booking_id, venue_event_type_id, venue_timeblock_id FROM event")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rental bookings to hold the data.
	var events []Event

	// Loop through the data and insert into the rental bookings slice.
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.BookingID, &event.VenueEventTypeID, &event.VenueTimeBlockID); err != nil {
			log.Fatalf("failed to scan: %v", err)
		}
		events = append(events, event)
	}

}

func GetEventsForVenue(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("failed to convert id to int: %v", err)
	}
	events, err := GetEventsForVenueId(idInt, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)

}

func CreateEvent(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Decode the request body into a RequestEvent struct.
	var details RequestEvent
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		log.Fatalf("failed to decode request: %v", err)
	}

	// Attempt to book the rental.
	eventID, err := AttemptToCreateEvent(details, db)
	if err != nil {
		log.Fatalf("failed to create event: %v", err)
	}

	if eventID == -1 {
		// Return a 409 Conflict if the rental is already booked.
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Venue is already booked"))

		return
	}

	if eventID == -2 {
		//minimum duration not met
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Minimum duration not met"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	// Return the rental booking ID as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventID)

}

func GetEventDetails(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	eventDetails, err := GetDetailsForEventID(eventId, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventDetails)

}
