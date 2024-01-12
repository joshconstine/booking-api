package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type BoatDefaultSettings struct {
	ID                      int
	BoatId                  int
	DailyCost               float32
	MinimunBookingDuration  int
	AdvertiseAtAllLocations bool
	fileId                  int
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

	boatDefaultSettings := []BoatDefaultSettings{
		{1, 1, 200.00, 1.00, true, 3},
		{2, 2, 200.00, 3.00, true, 3},
		{3, 3, 200.00, 2.00, true, 3},
		{4, 4, 200.00, 1.00, true, 3},
	}

	// Loop through the data and insert into the boat table
	for _, boatDefaultSetting := range boatDefaultSettings {
		insertQuery := "INSERT INTO boat_default_settings (boat_id, daily_cost, minimum_booking_duration, advertise_at_all_locations, file_id) VALUES (?,  ?, ?, ?, ?)"
		_, err := db.Exec(insertQuery, boatDefaultSetting.BoatId, boatDefaultSetting.DailyCost, boatDefaultSetting.MinimunBookingDuration, boatDefaultSetting.AdvertiseAtAllLocations, boatDefaultSetting.fileId)
		if err != nil {
			log.Fatal(err)

		}
	}
	fmt.Println("Inserted boat default settings into the boat_default_setting table")

	defer db.Close()

}
