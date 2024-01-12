package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type VenueEventType struct {
	ID          int
	VenueID     int
	EventTypeID int
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

	venueEventTypes := []VenueEventType{
		{1, 1, 1},
		{2, 1, 2},
		{3, 1, 3},
		{4, 2, 2},
		{5, 2, 3},
		{6, 1, 0},
		{7, 3, 3},
	}

	// Loop through the data and insert into the venue_event_type table
	for _, venueEventType := range venueEventTypes {
		insertQuery := "INSERT INTO venue_event_type (venue_id, event_type_id) VALUES (?, ?)"
		_, err := db.Exec(insertQuery, venueEventType.VenueID, venueEventType.EventTypeID)
		if err != nil {
			log.Fatal(err)

		}
	}
	log.Println("Inserted venue event types into the venue_event_type table")

	defer db.Close()

}
