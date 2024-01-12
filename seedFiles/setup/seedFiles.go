package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type File struct {
	ID   int
	Name string
	Url  string
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

	files := []File{
		{1, "Everett Resort Rental Agreement", "https://www.google.com"},
		{2, "Musky Inn Rental Agreement", "https://www.google.com"},
		{3, "Boat Rental Agreement", "https://www.google.com"},
	}

	// Insert the data into the files table
	for _, file := range files {
		_, err = db.Exec("INSERT INTO file (id, name, url) VALUES (?, ?, ?)", file.ID, file.Name, file.Url)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Successfully seeded file table")

}
