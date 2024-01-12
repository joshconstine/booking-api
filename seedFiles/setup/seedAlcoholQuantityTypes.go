package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type AlcoholQuantityType struct {
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

	alcoholQuantityTypes := []AlcoholQuantityType{
		{1, "750ml"},
		{2, "1000ml"},
		{3, "1750ml"},
		{4, "4pack cans"},
		{5, "6pack cans"},
		{6, "12pack cans"},
		{7, "18pack cans"},
		{8, "24pack cans"},
		{9, "30pack cans"},
		{10, "6pack bottles"},
		{11, "12pack bottles"},
		{12, "18pack bottles"},
		{13, "24pack bottles"},
		{14, "1/4 barrel"},
		{15, "1/2 barrel"},
	}

	// Insert the data into the alcohol_quantity_types table
	for _, alcoholQuantityType := range alcoholQuantityTypes {
		_, err = db.Exec("INSERT INTO alcohol_quantity_type (id, name) VALUES (?, ?)", alcoholQuantityType.ID, alcoholQuantityType.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded alcohol_quantity_type table")

	defer db.Close()

}
