package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

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

	//Seed Location table

	//Eagle River
	_, err = db.Exec("INSERT INTO location ( name) VALUES ( 'Eagle River')")
	if err != nil {
		log.Fatalf("failed to seed location table: %v", err)
	}
	//St Germain
	_, err = db.Exec("INSERT INTO location (name) VALUES ( 'St Germain')")
	if err != nil {
		log.Fatalf("failed to seed location table: %v", err)
	}

	//read from location table
	rows, err := db.Query("SELECT * FROM location")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	//log
	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		log.Println(id, name)
	}

}
