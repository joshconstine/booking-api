package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type AmenityType struct {
	ID   int
	Name string
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

	amenityTypes := []AmenityType{
		{1, "Kitchen"},
		{2, "Bathroom"},
		{3, "Laundry"},
		{4, "Entertainment"},
		{5, "Outdoor"},
		{6, "Utilities"},
		{7, "Safety"},
		{8, "Miscellaneous"},
	}

	for _, amenityType := range amenityTypes {
		_, err := db.Exec("INSERT INTO amenity_type (id, name) VALUES (?, ?)", amenityType.ID, amenityType.Name)
		if err != nil {
			log.Fatalf("failed to insert: %v", err)
		}
	}

	fmt.Println("seeded amenity_type")

	defer db.Close()
}
