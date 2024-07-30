package main

import (
	"booking-api/config"
	"booking-api/pkg/database"
	"fmt"
	"os"
)

func main() {

	// load config
	env, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	database.Connect(env.DSN)

	// create object storage client
	// objectStorage.CreateSession()

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		database.Migrate()

	}
	if cmd == "down" {
		// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		// 	log.Fatal(err)
		// }
	}
}
