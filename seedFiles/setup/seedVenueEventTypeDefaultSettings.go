package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type VenueEventTypeDefaultSettings struct {
	ID                     int
	VenueEventTypeID       int
	HourlyRate             float32
	MinimumBookingDuration int
	FlatFee                float32
	EarliestBookingTime    string
	LatestBookingTime      string
}

func main() {

	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("connected to PlanetScale")

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	venueEventTypeDefaultSettings := []VenueEventTypeDefaultSettings{
		{0, 1, 100.0, 3, 0.0, "15:00:00", "20:00:00"},
		{1, 2, 400.0, 1, 10.0, "15:00:00", "90:00:00"},
		{2, 3, 200.0, 2, 0.0, "15:00:00", "21:00:00"},
		{3, 4, 540.0, 0, 300.0, "15:00:00", "18:00:00"},
		{4, 5, 140.0, 0, 0.0, "15:00:00", "19:00:00"},
		{5, 6, 110.0, 2, 1.0, "15:00:00", "21:00:00"},
		{6, 7, 100.0, 1, 0.0, "15:00:00", "21:00:00"},
	}

	// Loop through the data and insert into the venue_event_type_default_settings table
	for _, venueEventTypeDefaultSetting := range venueEventTypeDefaultSettings {
		insertQuery := "INSERT INTO venue_event_type_default_settings (venue_event_type_id, hourly_rate, minimum_booking_duration, flat_fee, earliest_booking_time, latest_booking_time) VALUES (?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(insertQuery, venueEventTypeDefaultSetting.VenueEventTypeID, venueEventTypeDefaultSetting.HourlyRate, venueEventTypeDefaultSetting.MinimumBookingDuration, venueEventTypeDefaultSetting.FlatFee, venueEventTypeDefaultSetting.EarliestBookingTime, venueEventTypeDefaultSetting.LatestBookingTime)
		if err != nil {
			log.Fatal(err)

		}
	}

	log.Println("Inserted venue event type default settings into the venue_event_type_default_settings table")

	defer db.Close()
}
