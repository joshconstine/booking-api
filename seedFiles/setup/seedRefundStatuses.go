package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type RefundStatus struct {
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

	refundStatuses := []RefundStatus{
		{1, "Requested"},
		{2, "Dispursed"},
		{3, "Denied"},
	}
	// Insert the data into the refundStatus table
	for _, refundStatus := range refundStatuses {
		_, err = db.Exec("INSERT INTO refund_status (id, name) VALUES (?, ?)", refundStatus.ID, refundStatus.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded refund_statuses table")

	defer db.Close()

}
