package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RentalBedroom struct {
	ID            int
	RentalID      int
	Name          string
	Description   string
	Floor         int
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

	rentalBedrooms := []RentalBedroom{
		{1, 1, "Master Bedroom", "Spacious master bedroom with a king size bed and a private bathroom.", 1, nil},
		{2, 1, "Bedroom 2", "Cozy bedroom ", 1, nil},
		{3, 1, "Bedroom 3", "Cozy bedroom ", 1, nil},
		{4, 1, "Bedroom 4", "Cozy bedroom ", 1, nil},
		{5, 1, "Bedroom 5", "Cozy bedroom ", 1, nil},
		{6, 1, "Bedroom 6", "Cozy bedroom ", 1, nil},
		{7, 1, "Bedroom 7", "Cozy bedroom ", 2, nil},
		{8, 1, "Bedroom 8", "Cozy bedroom ", 2, nil},
		{9, 1, "Bedroom 9", "Cozy bedroom ", 2, nil},
		{10, 1, "Bedroom 10", "Cozy bedroom ", 2, nil},
		{11, 1, "Bedroom 11", "Cozy bedroom ", 2, nil},
		{12, 1, "Bedroom 12", "Cozy bedroom ", 2, nil},
		{13, 1, "Bedroom 13", "Cozy bedroom ", 1, nil},
		{14, 2, "The Nest", "Cozy bedroom", 1, nil},
		{15, 2, "The Den", "Cozy bedroom", 1, nil},
		{16, 3, "The Master", "Spacious master bedroom with a king size bed and a private bathroom.", 1, nil},
		{17, 3, "The Prince", "Cozy bedroom", 1, nil},
		{18, 3, "The Kids Room", "Cozy bedroom", 1, nil},
		{19, 3, "The Guest Room", "Cozy bedroom", 1, nil},
		{20, 3, "The Library", "Cozy bedroom", 1, nil},
		{21, 4, "The Master", "Spacious master bedroom with a king size bed and a private bathroom.", 1, nil},
		{22, 4, "The Prince", "Cozy bedroom", 1, nil},
		{23, 4, "The Kids Room", "Cozy bedroom", 1, nil},
		{24, 4, "The Guest Room", "Cozy bedroom", 1, nil},
		{25, 4, "The Attic", "Cozy bedroom", 2, nil},
		{26, 4, "The Timbers", "Cozy bedroom", 2, nil},
		{27, 4, "The Princess", "", 2, nil},
		{28, 6, "The Master", "Spacious master bedroom ", 1, nil},
		{29, 6, "The Theator", "Enjoy the big screen", 1, nil},
		{30, 6, "The Princess", "Cozy bedroom with attached bathroom", 1, nil},
		{31, 6, "The Prince", "Cozy bedroom With a view of the Lake. West Facing", 1, nil},
		{32, 6, "The Attic", "Cozy bedroom", 2, nil},
		{33, 6, "The Timbers", "Cozy bedroom with a lake view", 2, nil},
		{34, 6, "The Nest", "Cozy bedroom with a lake view", 2, nil},
		{35, 6, "The shelf", "Spacious bedroom", 2, nil},
		{36, 6, "The Library", "Cozy bedroom", 2, nil},
		{37, 6, "The Guest Room", "Cozy bedroom", 2, nil},
		{38, 6, "The Nursery", "Cozy bedroom", 2, nil},
		{39, 5, "The Master", "Cozy bedroom", 1, nil},
		{40, 5, "The Den", "Cozy bedroom", 1, nil},
		{41, 5, "The Kids Room", "Cozy bedroom", 1, nil},
		{42, 5, "The Guest Room", "Cozy bedroom", 1, nil},
		{43, 5, "The Library", "Cozy bedroom", 1, nil},
		// {44, 11, "The Nest", "Cozy bedroom", 2, nil},
		{44, 7, "The Nest", "Cozy bedroom", 2, nil},
	}

	// Insert the data into the rental_bedrooms table
	for _, rentalBedroom := range rentalBedrooms {
		_, err = db.Exec("INSERT INTO rental_bedroom (id, rental_id, name, description, floor, rental_photo_id) VALUES (?, ?, ?, ?, ?, ?)", rentalBedroom.ID, rentalBedroom.RentalID, rentalBedroom.Name, rentalBedroom.Description, rentalBedroom.Floor, rentalBedroom.RentalPhotoID)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data inserted into rental_bedrooms table successfully.")

	defer db.Close()
}
