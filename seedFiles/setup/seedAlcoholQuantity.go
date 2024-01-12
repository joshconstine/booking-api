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
	Price                 float32
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
		{1, 1, 13, 6.00},
		{2, 2, 9, 6.00},
		{3, 3, 5, 6.00},
		{4, 4, 5, 6.00},
		{5, 5, 5, 6.00},
		{6, 6, 5, 6.00},
		{7, 7, 5, 6.00},
		{8, 8, 5, 6.00},
		{9, 9, 5, 6.00},
		{10, 10, 5, 5.00},
		{11, 11, 5, 5.00},
		{12, 12, 5, 5.00},
		{13, 13, 5, 5.00},
		{14, 14, 1, 5.00},
		{15, 15, 1, 5.00},
		{16, 16, 1, 5.00},
		{17, 17, 2, 5.00},
		{18, 18, 2, 5.00},
		{19, 19, 2, 5.00},
		{20, 20, 2, 5.00},
		{21, 21, 2, 5.00},
		{22, 22, 2, 5.00},
		{23, 23, 2, 5.00},
		{24, 24, 2, 5.00},
		{25, 25, 2, 5.00},
		{26, 26, 2, 5.00},
		{27, 27, 2, 5.00},
		{28, 28, 2, 5.00},
		{29, 29, 2, 5.00},
		{30, 30, 2, 5.00},
		{31, 31, 2, 5.00},
		{32, 32, 2, 5.00},
		{33, 33, 2, 5.00},
		{34, 34, 2, 5.00},
		{35, 35, 2, 5.00},
		{36, 36, 2, 5.00},
		{37, 37, 2, 5.00},
		{38, 38, 2, 5.00},
		{39, 39, 2, 5.00},
		{40, 40, 2, 5.00},
		{41, 41, 6, 5.00},
		{42, 42, 6, 5.00},
		{43, 43, 6, 5.00},
		{44, 44, 6, 5.00},
		{45, 45, 6, 5.00},
		{46, 46, 2, 5.00},
		{47, 47, 2, 5.00},
	}

	// Insert the data into the alcohol_quantity table
	for _, alcoholQuantity := range alcoholQuantities {

		_, err = db.Exec("INSERT INTO alcohol_quantity (id, alcohol_id, alcohol_quantity_type_id, price) VALUES (?, ?, ?, ?)", alcoholQuantity.ID, alcoholQuantity.AlcoholID, alcoholQuantity.AlcoholQuantityTypeID, alcoholQuantity.Price)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded alcohol_quantity table")

	defer db.Close()

}
