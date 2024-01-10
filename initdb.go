package main

import (
    "database/sql"
    "log"
    "os"

    "github.com/joho/godotenv"
     _ "github.com/go-sql-driver/mysql"
)



func main() {

	//SQL CREATE TABLES
	

	//Rentals
	rentalsCreate := "CREATE TABLE IF NOT EXISTS rentals (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, location_id INT NOT NULL, bedrooms INT NOT NULL, bathrooms INT NOT NULL, PRIMARY KEY (id), KEY location_id (location_id))"
	rentalTimeblockCreate := "CREATE TABLE IF NOT EXISTS rental_timeblock (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, start_time DATETIME NOT NULL, end_time DATETIME NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id))"
	rentalUnitDefaultSettingsCreate := "CREATE TABLE IF NOT EXISTS rental_unit_default_settings (id INT NOT NULL AUTO_INCREMENT, rental_unit_id INT NOT NULL UNIQUE, nightly_cost DECIMAL(10, 2) NOT NULL, minimum_booking_duration INT NOT NULL, allows_pets BOOLEAN NOT NULL, cleaning_fee DECIMAL(10, 2) NOT NULL, check_in_time TIME NOT NULL, check_out_time TIME NOT NULL, PRIMARY KEY (id))"
	rentalUnitVariableSettingsCreate := "CREATE TABLE IF NOT EXISTS rental_unit_variable_settings (id INT NOT NULL AUTO_INCREMENT, rental_unit_id INT NOT NULL, start_date DATE NOT NULL, end_date DATE NOT NULL, minimum_booking_duration INT NOT NULL, nightly_cost DECIMAL(10, 2) NOT NULL, PRIMARY KEY (id), KEY rental_unit_id (rental_unit_id))"
	rentalPhotosCreate := "CREATE TABLE IF NOT EXISTS rental_photos (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, photo_id INT NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id), KEY photo_id (photo_id))"

	//Photos
	photosCreate := "CREATE TABLE IF NOT EXISTS photos (id INT NOT NULL AUTO_INCREMENT, url VARCHAR(255) NOT NULL, PRIMARY KEY (id))"

	

	//Locations
	locationsCreate := "CREATE TABLE IF NOT EXISTS locations (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"

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


	//Photos
	_, err = db.Exec(photosCreate)
	if err != nil {
		log.Fatalf("failed to create photos table: %v", err)
	}
	

	// Rentals

	_, err = db.Exec(rentalsCreate)
	if err != nil {
		log.Fatalf("failed to create rentals table: %v", err)
	}


	// Rental Settings
	_, err = db.Exec(rentalTimeblockCreate)
	if err != nil {
		log.Fatalf("failed to create rental_timeblock table: %v", err)
	}

	// Rental Unit Default Settings
	_, err = db.Exec(rentalUnitDefaultSettingsCreate)
	if err != nil {
		log.Fatalf("failed to create rental_unit_default_settings table: %v", err)
	}

	// Rental Unit Variable Settings
	_, err = db.Exec(rentalUnitVariableSettingsCreate)
	if err != nil {
		log.Fatalf("failed to create rental_unit_variable_settings table: %v", err)
	}

	// Rental Photos
	_, err = db.Exec(rentalPhotosCreate)
	if err != nil {
		log.Fatalf("failed to create rental_photos table: %v", err)
	}

	// Locations
	_, err = db.Exec(locationsCreate)
	if err != nil {
		log.Fatalf("failed to create locations table: %v", err)
	}
	




    rows, err := db.Query("SHOW TABLES")
    if err != nil {
        log.Fatalf("failed to query: %v", err)
    }
    defer rows.Close()

    var tableName string
    for rows.Next() {
        if err := rows.Scan(&tableName); err != nil {
            log.Fatalf("failed to scan row: %v", err)
        }
        log.Println(tableName)
		//describe each table
		describe := "DESCRIBE " + tableName
		rows2, err := db.Query(describe)
		if err != nil {
			log.Fatalf("failed to query: %v", err)

		}
		defer rows2.Close()

		
    }

    defer db.Close()
}
