package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RentalAmenity struct {
	ID        int
	RentalID  int
	AmenityID int
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

	rentalAmenities := []RentalAmenity{
		{1, 1, 1},
		{2, 1, 2},
		{3, 1, 3},
		{4, 1, 4},
		{5, 1, 5},
		{6, 1, 6},
		{7, 1, 7},
		{8, 1, 8},
		{9, 1, 15},
		{10, 1, 16},
		{11, 1, 17},
		{12, 1, 23},
		{13, 1, 24},
		{14, 1, 25},
		{15, 1, 26},
		{16, 1, 27},
		{17, 1, 28},
		{18, 1, 29},
		{19, 1, 33},
		{20, 1, 36},
		{21, 1, 37},
		{22, 1, 38},
		{23, 1, 39},
		{24, 1, 44},
		{25, 1, 45},
		{26, 2, 1},
		{27, 2, 2},
		{28, 2, 3},
		{29, 2, 4},
		{30, 2, 5},
		{31, 2, 6},
		{32, 2, 7},
		{33, 2, 8},
		{34, 2, 15},
		{35, 2, 16},
		{36, 2, 17},

		{39, 2, 25},
		{40, 2, 26},
		{41, 2, 27},
		{42, 2, 28},
		{43, 2, 29},
		{44, 2, 33},
		{45, 2, 36},
		{46, 2, 37},
		{47, 2, 38},
		{48, 2, 39},
		{49, 2, 44},
		{50, 2, 45},

		{51, 3, 1},
		{52, 3, 2},
		{53, 3, 3},
		{54, 3, 4},
		{55, 3, 5},
		{56, 3, 6},
		{57, 3, 7},
		{58, 3, 8},
		{59, 3, 15},
		{60, 3, 16},
		{61, 3, 17},

		{64, 3, 25},
		{65, 3, 26},
		{66, 3, 27},
		{67, 3, 28},
		{68, 4, 1},
		{69, 4, 2},
		{70, 4, 3},
		{71, 4, 4},
		{72, 4, 5},
		{73, 4, 6},
		{74, 4, 7},
		{75, 4, 8},
		{76, 4, 15},
		{77, 5, 1},
		{78, 5, 2},
		{79, 5, 3},
		{80, 5, 4},
		{81, 5, 5},
		{82, 5, 6},
		{83, 5, 7},
		{84, 5, 8},
		{85, 5, 15},
		{86, 5, 16},
		{87, 5, 17},
		{88, 6, 1},
		{89, 6, 2},
		{90, 6, 3},
		{91, 6, 4},
		{92, 6, 5},
		{93, 6, 6},
		{94, 6, 7},
		{95, 6, 8},
		{96, 6, 15},
		{97, 6, 16},
		{98, 6, 17},
		{99, 6, 23},
		{100, 6, 24},
		{101, 6, 25},
		{102, 6, 26},
		{103, 6, 27},
		{104, 6, 28},
		{105, 6, 29},
		{106, 6, 33},
		{107, 6, 36},
		{108, 6, 37},
		{109, 6, 38},
		{110, 6, 39},
		{111, 6, 44},
		{112, 6, 45},
		// {113, 11, 1},
		// {114, 11, 2},
		// {115, 11, 3},
		// {116, 11, 4},
		// {117, 11, 5},
		// {118, 11, 17},
		{113, 7, 1},
		{114, 7, 2},
		{115, 7, 3},
		{116, 7, 4},
		{117, 7, 5},
		{118, 7, 17},
	}

	// Insert the data into the rental_amenity table
	for _, ra := range rentalAmenities {
		_, err := db.Exec("INSERT INTO rental_amenity (rental_id, amenity_id) VALUES (?, ?)", ra.RentalID, ra.AmenityID)
		if err != nil {
			log.Fatalf("failed to insert rental amenity: %v", err)
		}
	}

	fmt.Println("successfully inserted rental amenities")

	defer db.Close()
}
