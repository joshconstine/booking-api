package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RentalPhoto struct {
	ID       int
	RentalID int
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

	rentalPhotos := []RentalPhoto{
		{1, 1, "lodge-cabin-eagle-river-09.PNG"},
		{2, 2, "morrey-cabin-eagle-river-04.PNG"},
		{3, 3, "gables-cabin-eagle-river-26.PNG"},
		{4, 4, "clubhouse-cabin-eagle-river-24.PNG"},
		{5, 5, "eisenhower-cabin-eagle-river-04.PNG"},
		{6, 6, "musky-inn-eagle-river-03.PNG"},
		{7, 7, "musky-inn-eagle-river-03.PNG"},
		{8, 8, "musky-inn-eagle-river-03.PNG"},
		{9, 9, "musky-inn-eagle-river-03.PNG"},
		{10, 10, "musky-inn-eagle-river-03.PNG"},
		{11, 11, "musky-inn-eagle-river-03.PNG"},
	}

	for _, rentalPhoto := range rentalPhotos {
		insertQuery := "INSERT INTO rental_photo (rental_id, photo_url) VALUES (?,  ?)"
		_, err := db.Exec(insertQuery, rentalPhoto.RentalID, rentalPhoto.PhotoURL)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("inserted rental photos")
}
