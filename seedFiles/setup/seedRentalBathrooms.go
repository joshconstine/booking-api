package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RentalBathroom struct {
	ID            int
	RentalID      int
	Name          string
	Description   string
	Floor         int
	bathtub       bool
	shower        bool
	RentalPhotoID *int
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

	rentalBathrooms := []RentalBathroom{
		{1, 1, " Bathroom 1", "Spacious bathroom.", 1, true, true, nil},
		{2, 1, "Bathroom 2", "Cozy bathroom.", 1, true, true, nil},
		{3, 1, "Bathroom 3", "Cozy bathroom.", 2, true, true, nil},
		{4, 1, "Bathroom 4", "Cozy bathroom.", 2, true, true, nil},
		{5, 1, "Bathroom 5", "Cozy bathroom.", 2, true, true, nil},
		{6, 2, "Main Bathroom", "Spacious bathroom.", 1, true, true, nil},
		{7, 3, "Bathroom 1", "Spacious bathroom.", 1, false, true, nil},
		{8, 3, "Bathroom 2", "Cozy bathroom.", 2, true, true, nil},
		{9, 3, "Bathroom 3", "Cozy bathroom.", 2, false, true, nil},
		{10, 4, "Bathroom 1", "Spacious bathroom.", 1, true, true, nil},
		{11, 4, "Bathroom 2", "Cozy bathroom.", 1, true, true, nil},
		{12, 5, "Bathroom 1", "Spacious bathroom.", 1, false, true, nil},
		{13, 5, "Bathroom 2", "Cozy bathroom.", 1, true, true, nil},
		{14, 6, "Wolf bathroom", "Spacious bathroom Near Kitchen", 1, true, true, nil},
		{15, 6, "Princess bathroom", "Cozy bathroom with washer/dryer", 1, false, true, nil},
		{16, 6, "Bedroom 3", "Spacious bathroom.", 1, true, true, nil},
		{17, 6, "Bedroom 4", "Cozy bathroom.", 1, true, true, nil},
		{18, 6, "Bedroom 5", "Cozy bathroom.", 2, true, true, nil},
		{19, 6, "Bedroom 6", "Cozy bathroom.", 2, true, true, nil},
		{20, 6, "Bedroom 7", "Cozy bathroom.", 2, true, true, nil},
		{21, 6, "Main Bathroom", "Spacious bathroom.", 1, true, true, nil},
	}

	for _, bathroom := range rentalBathrooms {
		_, err = db.Exec("INSERT INTO rental_bathroom (id, rental_id, name, description, floor, bathtub, shower, rental_photo_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", bathroom.ID, bathroom.RentalID, bathroom.Name, bathroom.Description, bathroom.Floor, bathroom.bathtub, bathroom.shower, bathroom.RentalPhotoID)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Inserted rental bathrooms")

	defer db.Close()

}
