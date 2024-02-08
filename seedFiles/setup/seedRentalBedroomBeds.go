package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RentalBedroomBed struct {
	ID              int
	RentalBedroomID int
	BedTypeID       int
}

func main() {

	//Load connection string from .env file
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

	rentalBedroomBeds := []RentalBedroomBed{

		{1, 1, 3},
		{2, 2, 3},
		{3, 3, 2},
		{4, 4, 2},
		{5, 5, 2},
		{6, 6, 2},
		{7, 7, 2},
		{8, 8, 2},
		{9, 9, 2},
		{10, 10, 2},
		{11, 11, 2},
		{12, 12, 2},
		{13, 13, 2},
		{14, 14, 3},
		{15, 15, 3},
		{16, 16, 3},
		{17, 17, 3},
		{18, 18, 3},
		{19, 19, 3},
		{20, 20, 3},
		{21, 21, 3},
		{22, 22, 3},
		{23, 23, 3},
		{24, 24, 2},
		{25, 25, 2},
		{26, 26, 2},
		{27, 27, 2},
		{28, 28, 2},
		{29, 29, 6},
		{30, 30, 2},
		{31, 31, 2},
		{32, 32, 2},
		{33, 33, 2},
		{34, 34, 2},
		{35, 35, 2},
		{36, 36, 2},
		{37, 37, 2},
		{38, 38, 2},
		{39, 39, 2},
		{40, 40, 2},
		{41, 41, 2},
		{42, 42, 2},
		{43, 43, 2},
		{44, 44, 2},
		{45, 29, 6},
		{46, 27, 7},
		{47, 28, 1},
	}

	// Insert the data into the rental_bedroom_beds table
	for _, rentalBedroomBed := range rentalBedroomBeds {
		_, err = db.Exec("INSERT INTO rental_bedroom_bed (id, rental_bedroom_id, bed_type_id) VALUES (?, ?, ?)", rentalBedroomBed.ID, rentalBedroomBed.RentalBedroomID, rentalBedroomBed.BedTypeID)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data inserted into rental_bedroom_bed table successfully.")
	// Close the connection to PlanetScale
	defer db.Close()

}
