package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type RentalUnitDefaultSettings struct {
	ID                     int
	RentalUnitID           int
	NightlyCost            float64
	MinimumBookingDuration int
	AllowsPets             bool
	CleaningFee            float64
	CheckInTime            string
	CheckOutTime           string
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

	// Data for rental_unit_default_settings
	defaultSettings := []RentalUnitDefaultSettings{
		{0, 1, 100.0, 2, true, 20.0, "14:00:00", "10:00:00"},
		{1, 2, 75.0, 3, false, 15.0, "15:00:00", "11:00:00"},
		// Add more data as needed
	}

	// Loop through the data and insert into rental_unit_default_settings table
	for _, settings := range defaultSettings {
		insertQuery := "INSERT INTO rental_unit_default_settings (rental_unit_id, nightly_cost, minimum_booking_duration, allows_pets, cleaning_fee, check_in_time, check_out_time) VALUES (?, ?, ?, ?, ?, ?, ?)"
		result, err := db.Exec(insertQuery, settings.RentalUnitID, settings.NightlyCost, settings.MinimumBookingDuration, settings.AllowsPets, settings.CleaningFee, settings.CheckInTime, settings.CheckOutTime)
		if err != nil {
			log.Fatal(err)
		}

		// Get the last inserted ID and update the struct
		lastInsertID, _ := result.LastInsertId()
		settings.ID = int(lastInsertID)
	}

	fmt.Println("Data inserted into rental_unit_default_settings table successfully.")

	


}
