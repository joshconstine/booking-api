package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type PaymentMethod struct {
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

	paymentMethods := []PaymentMethod{
		{1, "Cash"},
		{2, "PayPal"},
		{3, "Check"},
	}

	// Insert the data into the payment_methods table
	for _, paymentMethod := range paymentMethods {
		_, err = db.Exec("INSERT INTO payment_method (id, name) VALUES (?, ?)", paymentMethod.ID, paymentMethod.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded payment_method table")

	defer db.Close()

}
