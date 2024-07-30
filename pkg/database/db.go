package database

import (
	"booking-api/config"
	"database/sql"
	"log"
)

func ConnectToDB(env config.EnvVars) (*sql.DB, error) {
	db, err := sql.Open("mysql", env.DSN)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("connected to PlanetScale")

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	return db, nil

}
