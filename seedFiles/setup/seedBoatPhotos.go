package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type BoatPhoto struct {
	ID       int
	BoatID   int
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

	boatPhotos := []BoatPhoto{
		{1, 1, "boat1.jpeg"},
		{2, 2, "boat2.jpeg"},
		{3, 3, "boat4.jpeg"},
		{4, 4, "boat3.jpeg"},
	}

	for _, boatPhoto := range boatPhotos {
		insertQuery := "INSERT INTO boat_photo (boat_id, photo_url) VALUES (?,  ?)"
		_, err := db.Exec(insertQuery, boatPhoto.BoatID, boatPhoto.PhotoURL)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("inserted boat photos")
}
