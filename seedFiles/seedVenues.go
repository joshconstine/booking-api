package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Venue struct {
	ID         int
	Name       string
	LocationID int
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

	venues := []Venue{
		{1, "The SunRoom", 1},
		{2, "The Marina Bar", 1},
		{3, "The Musky Inn Bar", 2},
	}

	// Loop through the data and insert into the venue table
	for _, venue := range venues {
		insertQuery := "INSERT INTO venue (name, location_id) VALUES (?, ?)"
		_, err := db.Exec(insertQuery, venue.Name, venue.LocationID)
		if err != nil {
			log.Fatal(err)

		}
	}
	fmt.Println("Inserted venues into the venue table")

	//Retrieve the data from the venue table
	rows, err := db.Query("SELECT * FROM venue")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ID int
	var Name string
	var LocationID int

	for rows.Next() {
		err := rows.Scan(&ID, &Name, &LocationID)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(ID, Name, LocationID)
	}
}
