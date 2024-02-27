package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RentalUnitDefaultSettings struct {
	ID                     int
	RentalID               int
	NightlyCost            float64
	MinimumBookingDuration int
	AllowsPets             bool
	CleaningFee            float64
	CheckInTime            string
	CheckOutTime           string
	FileID                 int
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
		{0, 1, 75.0, 3, false, 15.0, "15:00:00", "11:00:00", 1},
		{1, 2, 75.0, 3, false, 15.0, "15:00:00", "11:00:00", 1},
		{2, 3, 125.0, 2, true, 25.0, "16:00:00", "12:00:00", 1},
		{3, 4, 150.0, 2, false, 30.0, "17:00:00", "13:00:00", 1},
		{4, 5, 200.0, 3, true, 35.0, "18:00:00", "14:00:00", 1},
		{5, 6, 250.0, 2, false, 40.0, "19:00:00", "15:00:00", 1},
		{6, 7, 300.0, 2, true, 45.0, "20:00:00", "16:00:00", 2},
		// {7, 8, 350.0, 3, false, 50.0, "21:00:00", "17:00:00", 2},
		// {8, 9, 400.0, 2, true, 55.0, "22:00:00", "18:00:00", 2},
		// {9, 10, 450.0, 2, false, 60.0, "23:00:00", "19:00:00", 2},
		// {10, 11, 500.0, 3, true, 65.0, "00:00:00", "20:00:00", 2},
		// Add more data as needed
	}

	// Loop through the data and insert into rental_unit_default_settings table
	for _, settings := range defaultSettings {
		insertQuery := "INSERT INTO rental_unit_default_settings (rental_id, nightly_cost, minimum_booking_duration, allows_pets, cleaning_fee, check_in_time, check_out_time, file_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
		result, err := db.Exec(insertQuery, settings.RentalID, settings.NightlyCost, settings.MinimumBookingDuration, settings.AllowsPets, settings.CleaningFee, settings.CheckInTime, settings.CheckOutTime, settings.FileID)
		if err != nil {
			log.Fatal(err)
		}

		// Get the last inserted ID and update the struct
		lastInsertID, _ := result.LastInsertId()
		settings.ID = int(lastInsertID)
	}

	fmt.Println("Data inserted into rental_unit_default_settings table successfully.")

}
