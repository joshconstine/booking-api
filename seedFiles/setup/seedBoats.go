package main

import (
	"booking-api/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

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

	boats := []models.Boat{
		{Name: "22' Pontoon", Occupancy: 10, MaxWeight: 2000},
		{Name: "24' Pontoon", Occupancy: 12, MaxWeight: 2500},
		{Name: "23' Ski Boat", Occupancy: 12, MaxWeight: 2500},
		{Name: "28' Pontoon", Occupancy: 8, MaxWeight: 1800},
	}

	// Loop through the data and insert into the boat table
	for _, boat := range boats {
		insertQuery := "INSERT INTO boats (name, occupancy, max_weight, created_at, updated_at, deleted_at) VALUES (?,  ?, ?, NOW(), NOW(), NULL)"
		_, err := db.Exec(insertQuery, boat.Name, boat.Occupancy, boat.MaxWeight)
		if err != nil {
			log.Fatal(err)

		}
	}
	fmt.Println("Inserted boats into the boat table")

	//Retrieve the data from the boat table
	rows, err := db.Query("SELECT id, name, occupancy, max_weight FROM boats")
	if err != nil {
		log.Fatal(err)
	}

	var boat models.Boat
	// Loop through the data and print the results to the console
	for rows.Next() {
		err := rows.Scan(&boat.ID, &boat.Name, &boat.Occupancy, &boat.MaxWeight)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Occupancy: %d, Max Weight: %d\n", boat.ID, boat.Name, boat.Occupancy, boat.MaxWeight)
	}

	defer db.Close()

}
