package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type EventType struct {
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

	eventTypes := []EventType{
		{1, "Wedding"},
		{2, "Wedding Reception"},
		{3, "Private Party"},
		{4, "other"},
	}

	// Insert the data into the event_types table
	for _, eventType := range eventTypes {
		_, err = db.Exec("INSERT INTO event_type (id, name) VALUES (?, ?)", eventType.ID, eventType.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded event_type table")
	defer db.Close()

}
