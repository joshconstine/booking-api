package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type AlcoholType struct {
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

	alcoholTypes := []AlcoholType{
		{1, "Beer"},
		{2, "IPA"},
		{3, "Gin"},
		{4, "Rum"},
		{5, "Tequila"},
		{6, "Vodka"},
		{7, "Whiskey"},
		{8, "Champagne"},
		{9, "Wine"},
		{10, "Soda"},
		{11, "Mixers"},
		{12, "Other"},
	}

	// Insert the data into the alcohol_types table
	for _, alcoholType := range alcoholTypes {
		_, err = db.Exec("INSERT INTO alcohol_type (id, name) VALUES (?, ?)", alcoholType.ID, alcoholType.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded alcohol_type table")

	defer db.Close()

}
