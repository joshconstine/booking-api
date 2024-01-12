package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type BookingStatus struct {
	ID   int
	Name string
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

	bookingStatuses := []BookingStatus{
		{1, "Requested"},
		{2, "Confirmed"},
		{3, "In Progress"},
		{4, "Completed"},
		{5, "Cancelled"},
	}

	// Insert the data into the booking_statuses table
	for _, bookingStatus := range bookingStatuses {
		_, err = db.Exec("INSERT INTO booking_status (id, name) VALUES (?, ?)", bookingStatus.ID, bookingStatus.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded booking_status table")
	defer db.Close()

}
