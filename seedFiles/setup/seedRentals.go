package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Rental struct {
	ID          int
	Name        string
	LocationID  int
	Bedrooms    int
	Bathrooms   int
	Description string
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

	rentals := []Rental{
		{1, "The Lodge", 1, 13, 5, "cozy up north cabin"},
		{2, "The Morey", 1, 2, 1, "cozy up north cabin"},
		{3, "The Gables", 1, 7, 3, "cozy up north cabin"},
		{4, "The Clubhouse", 1, 5, 2, "cozy up north cabin"},
		{5, "The Eisenhower", 1, 4, 2, "cozy up north cabin"},
		{6, "The Musky Inn", 2, 13, 7, "cozy up north cabin"},
		// {7, "The Musky Inn North", 2, 6, 4, "cozy up north cabin"},
		// {8, "The Musky Inn North + middle", 2, 9, 4, "cozy up north cabin"},
		// {9, "The Musky Inn South", 2, 4, 3, "cozy up north cabin"},
		// {10, "The Musky Inn South + middle", 2, 7, 5, "cozy up north cabin"},
		{7, "The Little Guy", 2, 1, 1, "cozy up north cabin"},
	}
	// Loop through the data and insert into the rentals table
	for _, rental := range rentals {
		insertQuery := "INSERT INTO rental (name, location_id, bedrooms, bathrooms, description) VALUES (?, ?, ?, ?, ?)"
		result, err := db.Exec(insertQuery, rental.Name, rental.LocationID, rental.Bedrooms, rental.Bathrooms, rental.Description)
		if err != nil {
			log.Fatal(err)
		}

		// Get the last inserted ID and update the struct
		lastInsertID, _ := result.LastInsertId()
		rental.ID = int(lastInsertID)
	}

	fmt.Println("Data inserted into rentals table successfully.")

	//Read from the rentals table

	//read from rentals table
	rows, err := db.Query("SELECT * FROM rental")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var ID int
	var Name string
	var LocationID int
	var Bedrooms int
	var Bathrooms int
	var Description string

	for rows.Next() {
		if err := rows.Scan(&ID, &Name, &LocationID, &Bedrooms, &Bathrooms, &Description); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		log.Println(ID, Name)
	}

}
