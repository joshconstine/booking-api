package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RentalStatus struct {
	ID       int
	RentalId int
	IsClean  bool
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

	rentalStatuses := []RentalStatus{
		{1, 1, false},
		{2, 2, true},
		{3, 3, false},
		{4, 4, true},
		{5, 5, false},
		{6, 6, true},
		{7, 7, false},
		{8, 8, true},
		{9, 9, false},
		{10, 10, true},
		{11, 11, false},
	}

	// Loop through the data and insert int
	for _, rentalStatus := range rentalStatuses {
		insertQuery := "INSERT INTO rental_status (rental_id, is_clean) VALUES (?,  ?)"
		_, err := db.Exec(insertQuery, rentalStatus.RentalId, rentalStatus.IsClean)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Inserted rental statuses into the rental_status table")

	// Add more data as needed

}
