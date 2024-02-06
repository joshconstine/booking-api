package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type VenuePhoto struct {
	ID       int
	VenueID  int
	PhotoURL string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	log.Println("connected to PlanetScale")

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	venuePhotos := []VenuePhoto{
		{1, 1, "IMG_1963.jpg"},
		{2, 2, "marina-bar-eagle-river-13.jpg"},
		{3, 3, "musky-inn-eagle-river-18.jpg"},
	}

	for _, venuePhoto := range venuePhotos {
		insertQuery := "INSERT INTO venue_photo (venue_id, photo_url) VALUES (?,  ?)"
		_, err := db.Exec(insertQuery, venuePhoto.VenueID, venuePhoto.PhotoURL)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("inserted venue photos")
}
