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
	rentalTimeblockCreate := "CREATE TABLE IF NOT EXISTS rental_timeblock (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, start_time DATETIME NOT NULL, end_time DATETIME NOT NULL, rental_booking_id INT, PRIMARY KEY (id), KEY rental_id (rental_id), KEY rental_booking_id (rental_booking_id))"
	rentalUnitDefaultSettingsCreate := "CREATE TABLE IF NOT EXISTS rental_unit_default_settings (id INT NOT NULL AUTO_INCREMENT, rental_unit_id INT NOT NULL UNIQUE, nightly_cost DECIMAL(10, 2) NOT NULL, minimum_booking_duration INT NOT NULL, allows_pets BOOLEAN NOT NULL, cleaning_fee DECIMAL(10, 2) NOT NULL, check_in_time TIME NOT NULL, check_out_time TIME NOT NULL, PRIMARY KEY (id))"
	rentalUnitVariableSettingsCreate := "CREATE TABLE IF NOT EXISTS rental_unit_variable_settings (id INT NOT NULL AUTO_INCREMENT, rental_unit_id INT NOT NULL, start_date DATE NOT NULL, end_date DATE NOT NULL, minimum_booking_duration INT NOT NULL, nightly_cost DECIMAL(10, 2) NOT NULL, PRIMARY KEY (id), KEY rental_unit_id (rental_unit_id))"
	rentalPhotosCreate := "CREATE TABLE IF NOT EXISTS rental_photos (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, photo_id INT NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id), KEY photo_id (photo_id))"
	rentalStatusCreate := "CREATE TABLE IF NOT EXISTS rental_status (id INT NOT NULL AUTO_INCREMENT, rental_unit_id INT NOT NULL UNIQUE, is_clean BOOLEAN, PRIMARY KEY (id))"

	//Boats
	boatsCreate := "CREATE TABLE IF NOT EXISTS boats (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, location_id INT NOT NULL, occupancy INT NOT NULL, max_weight INT NOT NULL, PRIMARY KEY (id), KEY location_id (location_id))"
	boatTimeblockCreate := "CREATE TABLE IF NOT EXISTS boat_timeblock (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL, start_time DATETIME NOT NULL, end_time DATETIME NOT NULL, boat_booking_id INT, PRIMARY KEY (id), KEY boat_id (boat_id), KEY boat_booking_id (boat_booking_id))"
	boatPhotosCreate := "CREATE TABLE IF NOT EXISTS boat_photos (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL, photo_id INT NOT NULL, PRIMARY KEY (id), KEY boat_id (boat_id), KEY photo_id (photo_id))"
	boatStatusCreate := "CREATE TABLE IF NOT EXISTS boat_status (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL UNIQUE, is_clean BOOLEAN, lowFuel BOOLEAN, current_location_id INT NOT NULL, PRIMARY KEY (id), KEY current_location_id (current_location_id))"
	boatDefaultSettingsCreate := "CREATE TABLE IF NOT EXISTS boat_default_settings (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL UNIQUE, daily_cost DECIMAL(10, 2) NOT NULL, minimum_booking_duration INT NOT NULL, advertist_all_locations BOOLEAN NOT NULL, PRIMARY KEY (id))"


	//Photos
	photosCreate := "CREATE TABLE IF NOT EXISTS photos (id INT NOT NULL AUTO_INCREMENT, url VARCHAR(255) NOT NULL, PRIMARY KEY (id))"

	//Users
	usersCreate := "CREATE TABLE IF NOT EXISTS users (id INT NOT NULL AUTO_INCREMENT, email VARCHAR(255) NOT NULL, firstName VARCHAR(255) NOT NULL, lastName VARCHAR(255) NOT NULL, phoneNumber VARCHAR(15) NOT NULL, PRIMARY KEY (id))"



	//Locations
	locationsCreate := "CREATE TABLE IF NOT EXISTS locations (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"

	//Bookings
	bookingStatusesCreate := "CREATE TABLE IF NOT EXISTS booking_statuses (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	bookingsCreate := "CREATE TABLE IF NOT EXISTS bookings (id INT NOT NULL AUTO_INCREMENT, user_id INT NOT NULL, booking_status_id INT NOT NULL, booking_details_id INT NOT NULL, PRIMARY KEY (id), KEY user_id (user_id), KEY booking_status_id (booking_status_id), KEY booking_details_id (booking_details_id))"
	bookingDetailsCreate := "CREATE TABLE IF NOT EXISTS booking_details (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL UNIQUE, payment_complete BOOLEAN NOT NULL, payment_due_date DATE NOT NULL, documents_signed BOOLEAN NOT NULL, PRIMARY KEY (id))"
	rentalBookingCreate := "CREATE TABLE IF NOT EXISTS rental_booking (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, booking_id INT NOT NULL, rental_time_block_id INT NOT NULL, booking_status_id INT NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id), KEY booking_id (booking_id), KEY rental_time_block_id (rental_time_block_id), KEY booking_status_id (booking_status_id))"
	rentalBookingCostsCreate := "CREATE TABLE IF NOT EXISTS rental_booking_costs (id INT NOT NULL AUTO_INCREMENT, rental_booking_id INT NOT NULL, booking_cost_items_id INT NOT NULL, PRIMARY KEY (id), KEY rental_booking_id (rental_booking_id), KEY booking_cost_items_id (booking_cost_items_id))"
	boatBookingCreate := "CREATE TABLE IF NOT EXISTS boat_booking (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL, booking_id INT NOT NULL, boat_time_block_id INT NOT NULL, booking_status_id INT NOT NULL, PRIMARY KEY (id), KEY boat_id (boat_id), KEY booking_id (booking_id), KEY boat_time_block_id (boat_time_block_id), KEY booking_status_id (booking_status_id))"
	boatBookingCostsCreate := "CREATE TABLE IF NOT EXISTS boat_booking_costs (id INT NOT NULL AUTO_INCREMENT, boat_booking_id INT NOT NULL, booking_cost_items_id INT NOT NULL, PRIMARY KEY (id), KEY boat_booking_id (boat_booking_id), KEY booking_cost_items_id (booking_cost_items_id))"

	bookingPaymentCreate := "CREATE TABLE IF NOT EXISTS booking_payment (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, payment_amount DECIMAL(10, 2) NOT NULL, paypal_order_id INT, payment_method_id INT NOT NULL, PRIMARY KEY (id), KEY booking_id (booking_id), KEY payment_method_id (payment_method_id))"
	paymentMethodCreate := "CREATE TABLE IF NOT EXISTS payment_method (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"

	bookingCostTypesCreate := "CREATE TABLE IF NOT EXISTS booking_cost_types (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	bookingCostItemsCreate := "CREATE TABLE IF NOT EXISTS booking_cost_items (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, booking_cost_type_id INT NOT NULL, ammount DECIMAL(10, 2) NOT NULL, PRIMARY KEY (id), KEY booking_id (booking_id), KEY booking_cost_type_id (booking_cost_type_id))"

	//refunds

	refundStatusesCreate := "CREATE TABLE IF NOT EXISTS refund_statuses (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	refundRequestsCreate := "CREATE TABLE IF NOT EXISTS refund_requests (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, refund_status_id INT NOT NULL, refund_amount DECIMAL(10, 2) NOT NULL, PRIMARY KEY (id), KEY booking_id (booking_id), KEY refund_status_id (refund_status_id))"



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

	// Rental  Status
	_, err = db.Exec(rentalStatusCreate)
	if err != nil {
		log.Fatalf("failed to create rental_unit_status table: %v", err)
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

	// Boats
	_, err = db.Exec(boatsCreate)
	if err != nil {
		log.Fatalf("failed to create boats table: %v", err)
	}

	// Boat Timeblocks
	_, err = db.Exec(boatTimeblockCreate)
	if err != nil {
		log.Fatalf("failed to create boat_timeblock table: %v", err)
	}
	// Boat Status
	_, err = db.Exec(boatStatusCreate)
	if err != nil {
		log.Fatalf("failed to create boat_status table: %v", err)
	}

	// Boat Default Settings
	_, err = db.Exec(boatDefaultSettingsCreate)
	if err != nil {
		log.Fatalf("failed to create boat_default_settings table: %v", err)
	}


	// Boat Photos
	_, err = db.Exec(boatPhotosCreate)
	if err != nil {
		log.Fatalf("failed to create boat_photos table: %v", err)
	}




	// Locations
	_, err = db.Exec(locationsCreate)
	if err != nil {
		log.Fatalf("failed to create locations table: %v", err)
	}
	
	// Users
	_, err = db.Exec(usersCreate)
	if err != nil {
		log.Fatalf("failed to create users table: %v", err)
	}

	//Booking Statuses
	_, err = db.Exec(bookingStatusesCreate)
	if err != nil {
		log.Fatalf("failed to create booking_statuses table: %v", err)
	}

	//Bookings
	_, err = db.Exec(bookingsCreate)
	if err != nil {
		log.Fatalf("failed to create bookings table: %v", err)
	}

	//Booking Details
	_, err = db.Exec(bookingDetailsCreate)
	if err != nil {
		log.Fatalf("failed to create booking_details table: %v", err)
	}

	//Rental Booking
	_, err = db.Exec(rentalBookingCreate)
	if err != nil {
		log.Fatalf("failed to create rental_booking table: %v", err)
	}

	//Rental Booking Costs
	_, err = db.Exec(rentalBookingCostsCreate)
	if err != nil {
		log.Fatalf("failed to create rental_booking_costs table: %v", err)
	}

	//Boat Booking
	_, err = db.Exec(boatBookingCreate)
	if err != nil {
		log.Fatalf("failed to create boat_booking table: %v", err)
	}

	//Boat Booking Costs
	_, err = db.Exec(boatBookingCostsCreate)
	if err != nil {
		log.Fatalf("failed to create boat_booking_costs table: %v", err)
	}


	//Booking Cost Types
	_, err = db.Exec(bookingCostTypesCreate)
	if err != nil {
		log.Fatalf("failed to create booking_cost_types table: %v", err)
	}

	//Booking Cost Items
	_, err = db.Exec(bookingCostItemsCreate)
	if err != nil {
		log.Fatalf("failed to create booking_cost_items table: %v", err)
	}

	//Payment Method
	_, err = db.Exec(paymentMethodCreate)
	if err != nil {
		log.Fatalf("failed to create payment_method table: %v", err)
	}
	

	//Booking Payment
	_, err = db.Exec(bookingPaymentCreate)
	if err != nil {
		log.Fatalf("failed to create booking_payment table: %v", err)
	}

	//Refund Statuses
	_, err = db.Exec(refundStatusesCreate)
	if err != nil {
		log.Fatalf("failed to create refund_statuses table: %v", err)
	}

	//Refund Requests
	_, err = db.Exec(refundRequestsCreate)
	if err != nil {
		log.Fatalf("failed to create refund_requests table: %v", err)
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
