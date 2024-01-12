package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Boat struct {
	ID        int
	Name      string
	occupancy int
	maxWeight int
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

	boats := []Boat{
		{1, "22' Pontoon", 10, 2000},
		{2, "24' Pontoon", 12, 2500},
		{3, "23' Ski Boat", 12, 2500},
		{4, "28' Pontoon", 8, 1800},
	}

	// Loop through the data and insert into the boat table
	for _, boat := range boats {
		insertQuery := "INSERT INTO boat (name, occupancy, max_weight) VALUES (?,  ?, ?)"
		_, err := db.Exec(insertQuery, boat.Name, boat.occupancy, boat.maxWeight)
		if err != nil {
			log.Fatal(err)

		}
	}
	fmt.Println("Inserted boats into the boat table")

	//Retrieve the data from the boat table
	rows, err := db.Query("SELECT * FROM boat")
	if err != nil {
		log.Fatal(err)
	}

	var boat Boat
	// Loop through the data and print the results to the console
	for rows.Next() {
		err := rows.Scan(&boat.ID, &boat.Name, &boat.occupancy, &boat.maxWeight)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Occupancy: %d, Max Weight: %d\n", boat.ID, boat.Name, boat.occupancy, boat.maxWeight)
	}

}
