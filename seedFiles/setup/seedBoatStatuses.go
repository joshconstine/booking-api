package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type BoatStatus struct {
	ID                int
	BoatId            int
	IsClean           bool
	LowFuel           bool
	CurrentLocationID int
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

	boatStatuses := []BoatStatus{
		{1, 1, false, true, 1},
		{2, 2, true, true, 1},
		{3, 3, false, false, 1},
		{4, 4, true, false, 2},
	}

	// Loop through the data and insert into the boat table
	for _, boatStatus := range boatStatuses {
		insertQuery := "INSERT INTO boat_status (boat_id, is_clean, low_fuel, current_location_id) VALUES (?,  ?, ?, ?)"
		_, err := db.Exec(insertQuery, boatStatus.BoatId, boatStatus.IsClean, boatStatus.LowFuel, boatStatus.CurrentLocationID)
		if err != nil {
			log.Fatal(err)

		}
	}
	log.Println("Inserted boat statuses into the boat_status table")

}
