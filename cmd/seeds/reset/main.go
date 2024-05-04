package main

import (
	"booking-api/config"
	"booking-api/pkg/database"
	"fmt"
	"log"
	"os"
)

func main() {

	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()
	// load config
	env, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	database.Connect(env.DSN)

	tables := []string{
		"bookings",
		"booking_cost_items",
		"booking_documents",
		"booking_payments",
		"entity_bookings",
		"users",
		"chats",
		"chat_messages",
		"rentals",
		"boats",
		"accounts",
		"entity_timeblocks",
		"locations",
		"entity_bookings",
		"entity_booking_rules",
		"entity_booking_costs",
		// "tax_rates",
	}

	for _, table := range tables {
		result := database.Instance.Exec("DELETE FROM " + table)
		if result.Error != nil {
			log.Fatalf("failed to truncate table %s: %v", table, err)
		}
	}

	// SeedBoookingWithConflicts(database.Instance)
	log.Println("Database reset  Completed!")

}
