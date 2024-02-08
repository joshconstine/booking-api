package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Amenity struct {
	ID            int
	AmenityTypeID int
	Name          string
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

	amenities := []Amenity{
		{1, 1, "Refrigerator"},
		{2, 1, "Microwave"},
		{3, 1, "Oven"},
		{4, 1, "Stove"},
		{5, 1, "Dishwasher"},
		{6, 1, "Coffee Maker"},
		{7, 1, "Toaster"},
		{8, 1, "Blender"},
		{9, 1, "Food Processor"},
		{10, 1, "Slow Cooker"},
		{11, 1, "Stand Mixer"},
		{12, 1, "Waffle Iron"},
		{13, 1, "Rice Cooker"},
		{14, 1, "Electric Kettle"},
		{15, 2, "Hair Dryer"},
		{16, 2, "Cleaning Supplies"},
		{17, 2, "Toilet Paper"},
		{18, 2, "Shampoo"},
		{19, 2, "Conditioner"},
		{20, 2, "Body Wash"},
		{21, 2, "Hand Soap"},
		{22, 2, "Towels"},
		{23, 3, "Washer"},
		{24, 3, "Dryer"},
		{25, 3, "Iron"},
		{26, 3, "Ironing Board"},
		{27, 4, "TV"},
		{28, 4, "Cable"},
		{29, 4, "Netflix"},
		{30, 4, "Hulu"},
		{31, 4, "Amazon Prime"},
		{32, 4, "Apple TV"},
		{33, 6, "WiFi"},
		{34, 5, "Patio"},
		{35, 5, "Balcony"},
		{36, 5, "Grill"},
		{37, 5, "Fire Pit"},
		{38, 6, "Central Air Conditioning"},
		{39, 6, "Central Heating"},
		{40, 6, "Fan"},
		{41, 6, "Space Heater"},
		{42, 7, "Smoke Detector"},
		{43, 7, "Carbon Monoxide Detector"},
		{44, 7, "First Aid Kit"},
		{45, 7, "Fire Extinguisher"},
		{46, 8, "Luggage Dropoff Allowed"},
		{47, 8, "Long Term Stays Allowed"},
		{48, 8, "Private Entrance"},
	}

	for _, amenity := range amenities {
		_, err := db.Exec("INSERT INTO amenity (id, amenity_type_id, name) VALUES (?, ?, ?)", amenity.ID, amenity.AmenityTypeID, amenity.Name)
		if err != nil {
			log.Fatalf("failed to insert: %v", err)
		}
	}

	fmt.Println("seeded amenity table successfully.")

	defer db.Close()

}
