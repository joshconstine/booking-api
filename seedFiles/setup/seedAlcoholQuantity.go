package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type AlcoholQuantity struct {
	ID                    int
	AlcoholID             int
	AlcoholQuantityTypeID int
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

	alcoholQuantities := []AlcoholQuantity{
		{1, 1, 13},
		{2, 2, 9},
		{3, 3, 5},
		{4, 4, 5},
		{5, 5, 5},
		{6, 6, 5},
		{7, 7, 5},
		{8, 8, 5},
		{9, 9, 5},
		{10, 10, 5},
		{11, 11, 5},
		{12, 12, 5},
		{13, 13, 5},
		{14, 14, 1},
		{15, 15, 1},
		{16, 16, 1},
		{17, 17, 2},
		{18, 18, 2},
		{19, 19, 2},
		{20, 20, 2},
		{21, 21, 2},
		{22, 22, 2},
		{23, 23, 2},
		{24, 24, 2},
		{25, 25, 2},
		{26, 26, 2},
		{27, 27, 2},
		{28, 28, 2},
		{29, 29, 2},
		{30, 30, 2},
		{31, 31, 2},
		{32, 32, 2},
		{33, 33, 2},
		{34, 34, 2},
		{35, 35, 2},
		{36, 36, 2},
		{37, 37, 2},
		{38, 38, 2},
		{39, 39, 2},
		{40, 40, 2},
		{41, 41, 6},
		{42, 42, 6},
		{43, 43, 6},
		{44, 44, 6},
		{45, 45, 6},
		{46, 46, 2},
		{47, 47, 2},
	}

	// Insert the data into the alcohol_quantity table
	for _, alcoholQuantity := range alcoholQuantities {
		_, err = db.Exec("INSERT INTO alcohol_quantity (id, alcohol_id, alcohol_quantity_type_id) VALUES (?, ?, ?)", alcoholQuantity.ID, alcoholQuantity.AlcoholID, alcoholQuantity.AlcoholQuantityTypeID)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded alcohol_quantity table")

	defer db.Close()

}
