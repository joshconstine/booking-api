package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type BookingCostType struct {
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

	bookingCostTypes := []BookingCostType{
		{1, "Tax"},
		{2, "Cleaning Fee"},
		{3, "Cabin Rental Cost"},
		{4, "Boat Rental Cost"},
		{5, "Gas Refil fee"},
		{6, "Labor"},
		{7, "Damage Fee"},
		{8, "Wedding Fee"},
		{9, "Event fee"},
		{10, "Event Fee Flat"},
		{11, "Open Bar Fee"},
		{12, "Cancelation Fee"},
		{13, "Alcohol"},
	}

	// Insert the data into the booking_cost_types table
	for _, bookingCostType := range bookingCostTypes {
		_, err = db.Exec("INSERT INTO booking_cost_type (id, name) VALUES (?, ?)", bookingCostType.ID, bookingCostType.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded booking_cost_type table")
	defer db.Close()

}
