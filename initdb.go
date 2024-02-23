package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() {

	//SQL CREATE TABLES

	//Rentals
	rentalCreate := "CREATE TABLE IF NOT EXISTS rental (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, location_id INT NOT NULL, bedrooms INT NOT NULL, bathrooms INT NOT NULL, description VARCHAR(255), PRIMARY KEY (id), KEY location_id (location_id))"
	rentalTimeblockCreate := "CREATE TABLE IF NOT EXISTS rental_timeblock (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, start_time DATETIME NOT NULL, end_time DATETIME NOT NULL, rental_booking_id INT, PRIMARY KEY (id), KEY rental_id (rental_id), KEY rental_booking_id (rental_booking_id))"
	rentalUnitDefaultSettingsCreate := "CREATE TABLE IF NOT EXISTS rental_unit_default_settings (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL UNIQUE, nightly_cost DECIMAL(10, 2) NOT NULL, minimum_booking_duration INT NOT NULL, allows_pets BOOLEAN NOT NULL, cleaning_fee DECIMAL(10, 2) NOT NULL, check_in_time TIME NOT NULL, check_out_time TIME NOT NULL, file_id INT NOT NULL, PRIMARY KEY (id), KEY file_id (file_id))"
	rentalUnitVariableSettingsCreate := "CREATE TABLE IF NOT EXISTS rental_unit_variable_settings (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, start_date DATE NOT NULL, end_date DATE NOT NULL, minimum_booking_duration INT NOT NULL, nightly_cost DECIMAL(10, 2) NOT NULL, cleaning_fee DECIMAL(10, 2) NOT NULL, event_required BOOLEAN NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id))"
	rentalPhotoCreate := "CREATE TABLE IF NOT EXISTS rental_photo (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, photo_url VARCHAR(255) NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id))"
	rentalStatusCreate := "CREATE TABLE IF NOT EXISTS rental_status (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL UNIQUE, is_clean BOOLEAN, PRIMARY KEY (id))"

	//rental Bedrooms
	bedTypeCreate := "CREATE TABLE IF NOT EXISTS bed_type (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"

	rentalBedroomCreate := "CREATE TABLE IF NOT EXISTS rental_bedroom (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, rental_photo_id INT, name VARCHAR(255), description VARCHAR(255), floor INT NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id), KEY rental_photo_id (rental_photo_id))"
	rentalBedroomBedCreate := "CREATE TABLE IF NOT EXISTS rental_bedroom_bed (id INT NOT NULL AUTO_INCREMENT, rental_bedroom_id INT NOT NULL, bed_type_id INT NOT NULL, PRIMARY KEY (id), KEY rental_bedroom_id (rental_bedroom_id), KEY bed_type_id (bed_type_id))"

	//rental Bathrooms
	rentallBathroomCreate := "CREATE TABLE IF NOT EXISTS rental_bathroom (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, rental_photo_id INT, name VARCHAR(255), description VARCHAR(255), floor INT NOT NULL, shower BOOLEAN NOT NULL, bathtub BOOLEAN NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id), KEY rental_photo_id (rental_photo_id))"

	//ameniy
	amenityTypeCreate := "CREATE TABLE IF NOT EXISTS amenity_type (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	amenityCreate := "CREATE TABLE IF NOT EXISTS amenity (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, amenity_type_id INT NOT NULL, PRIMARY KEY (id), KEY amenity_type_id (amenity_type_id))"

	rentalAmenityCreate := "CREATE TABLE IF NOT EXISTS rental_amenity(id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, amenity_id INT NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id), KEY amenity_id (amenity_id))"

	//Boats
	boatCreate := "CREATE TABLE IF NOT EXISTS boat (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, occupancy INT NOT NULL, max_weight INT NOT NULL, PRIMARY KEY (id))"
	boatTimeblockCreate := "CREATE TABLE IF NOT EXISTS boat_timeblock (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL, start_time DATETIME NOT NULL, end_time DATETIME NOT NULL, boat_booking_id INT, PRIMARY KEY (id), KEY boat_id (boat_id), KEY boat_booking_id (boat_booking_id))"
	boatPhotoCreate := "CREATE TABLE IF NOT EXISTS boat_photo (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL, photo_url VARCHAR(255) NOT NULL, PRIMARY KEY (id), KEY boat_id (boat_id))"
	boatStatusCreate := "CREATE TABLE IF NOT EXISTS boat_status (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL UNIQUE, is_clean BOOLEAN, low_fuel BOOLEAN, current_location_id INT NOT NULL, PRIMARY KEY (id), KEY current_location_id (current_location_id))"
	boatDefaultSettingsCreate := "CREATE TABLE IF NOT EXISTS boat_default_settings (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL UNIQUE, daily_cost DECIMAL(10, 2) NOT NULL, minimum_booking_duration INT NOT NULL, advertise_at_all_locations BOOLEAN NOT NULL, file_id INT NOT NULL, PRIMARY KEY (id), KEY file_id (file_id))"
	boatVariableSettingsCreate := "CREATE TABLE IF NOT EXISTS boat_variable_settings (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL, start_date DATE NOT NULL, end_date DATE NOT NULL, daily_cost DECIMAL(10, 2) NOT NULL, minimum_booking_duration INT NOT NULL, PRIMARY KEY (id), KEY boat_id (boat_id))"

	//Photo
	photoCreate := "CREATE TABLE IF NOT EXISTS photo (id INT NOT NULL AUTO_INCREMENT, url VARCHAR(255) NOT NULL, PRIMARY KEY (id))"

	//User
	userCreate := "CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, email VARCHAR(255) NOT NULL, first_name VARCHAR(255) NOT NULL, last_name VARCHAR(255) NOT NULL, phone_number VARCHAR(15) NOT NULL, PRIMARY KEY (id))"

	//Location
	locationCreate := "CREATE TABLE IF NOT EXISTS location (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"

	//Bookings
	bookingStatusCreate := "CREATE TABLE IF NOT EXISTS booking_status (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	bookingCreate := "CREATE TABLE IF NOT EXISTS booking (id INT NOT NULL AUTO_INCREMENT, user_id INT NOT NULL, booking_status_id INT NOT NULL, booking_details_id INT NOT NULL, PRIMARY KEY (id), KEY user_id (user_id), KEY booking_status_id (booking_status_id), KEY booking_details_id (booking_details_id))"
	bookingDetailsCreate := "CREATE TABLE IF NOT EXISTS booking_details (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL UNIQUE, payment_complete BOOLEAN NOT NULL, payment_due_date DATE NOT NULL, documents_signed BOOLEAN NOT NULL, booking_start_date DATETIME NOT NULL, invoice_id VARCHAR(255), PRIMARY KEY (id))"
	rentalBookingCreate := "CREATE TABLE IF NOT EXISTS rental_booking (id INT NOT NULL AUTO_INCREMENT, rental_id INT NOT NULL, booking_id INT NOT NULL, rental_time_block_id INT NOT NULL, booking_status_id INT NOT NULL, booking_file_id INT NOT NULL, PRIMARY KEY (id), KEY rental_id (rental_id), KEY booking_id (booking_id), KEY rental_time_block_id (rental_time_block_id), KEY booking_status_id (booking_status_id), KEY booking_file_id (booking_file_id))"
	rentalBookingCostCreate := "CREATE TABLE IF NOT EXISTS rental_booking_cost (id INT NOT NULL AUTO_INCREMENT, rental_booking_id INT NOT NULL, booking_cost_item_id INT NOT NULL, PRIMARY KEY (id), KEY rental_booking_id (rental_booking_id), KEY booking_cost_item_id (booking_cost_item_id))"
	boatBookingCreate := "CREATE TABLE IF NOT EXISTS boat_booking (id INT NOT NULL AUTO_INCREMENT, boat_id INT NOT NULL, booking_id INT NOT NULL, boat_time_block_id INT NOT NULL, booking_status_id INT NOT NULL, location_id INT NOT NULL, booking_file_id INT NOT NULL, PRIMARY KEY (id), KEY boat_id (boat_id), KEY booking_id (booking_id), KEY boat_time_block_id (boat_time_block_id), KEY booking_status_id (booking_status_id), KEY location_id (location_id), KEY booking_file_id (booking_file_id))"
	boatBookingCostCreate := "CREATE TABLE IF NOT EXISTS boat_booking_cost (id INT NOT NULL AUTO_INCREMENT, boat_booking_id INT NOT NULL, booking_cost_item_id INT NOT NULL, PRIMARY KEY (id), KEY boat_booking_id (boat_booking_id), KEY booking_cost_item_id (booking_cost_item_id))"

	bookingPaymentCreate := "CREATE TABLE IF NOT EXISTS booking_payment (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, payment_amount DECIMAL(10, 2) NOT NULL, paypal_order_id INT, payment_method_id INT NOT NULL,  payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP , PRIMARY KEY (id), KEY booking_id (booking_id), KEY payment_method_id (payment_method_id))"
	paymentMethodCreate := "CREATE TABLE IF NOT EXISTS payment_method (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"

	bookingCostTypeCreate := "CREATE TABLE IF NOT EXISTS booking_cost_type (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	bookingCostItemCreate := "CREATE TABLE IF NOT EXISTS booking_cost_item (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, booking_cost_type_id INT NOT NULL, amount DECIMAL(10, 2) NOT NULL, PRIMARY KEY (id), KEY booking_id (booking_id), KEY booking_cost_type_id (booking_cost_type_id))"

	alcoholOrderCreate := "CREATE TABLE IF NOT EXISTS alcohol_order (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, PRIMARY KEY (id))"
	alcoholOrderItemCreate := "CREATE TABLE IF NOT EXISTS alcohol_order_item (id INT NOT NULL AUTO_INCREMENT, alcohol_order_id INT NOT NULL,alcohol_order_booking_cost_id INT NOT NULL,alcohol_quantity_id INT NOT NULL,quantity INT NOT NULL, PRIMARY KEY (id), KEY alcohol_order_id (alcohol_order_id), KEY alcohol_order_booking_cost_id (alcohol_order_booking_cost_id), KEY alcohol_quantity_id (alcohol_quantity_id))"
	alcoholOrderBookingCost := "CREATE TABLE IF NOT EXISTS alcohol_order_booking_cost (id INT NOT NULL AUTO_INCREMENT, alcohol_order_id INT NOT NULL,alcohol_order_item_id INT NOT NULL, booking_cost_item_id INT NOT NULL,quantity INT NOT NULL, PRIMARY KEY (id), KEY alcohol_order_id (alcohol_order_id), KEY alcohol_order_item_id (alcohol_order_item_id),  KEY booking_cost_item_id (booking_cost_item_id))"

	bookingFileCreate := "CREATE TABLE IF NOT EXISTS booking_file (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, file_id INT NOT NULL, PRIMARY KEY (id), KEY booking_id (booking_id), KEY file_id (file_id))"

	//refunds

	refundStatusCreate := "CREATE TABLE IF NOT EXISTS refund_status (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	refundRequestCreate := "CREATE TABLE IF NOT EXISTS refund_request (id INT NOT NULL AUTO_INCREMENT, booking_id INT NOT NULL, refund_status_id INT NOT NULL, refund_amount DECIMAL(10, 2) NOT NULL, PRIMARY KEY (id), KEY booking_id (booking_id), KEY refund_status_id (refund_status_id))"

	//Alcohol
	alcoholTypeCreate := "CREATE TABLE IF NOT EXISTS alcohol_type (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	alcoholQuantityTypeCreate := "CREATE TABLE IF NOT EXISTS alcohol_quantity_type (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	alcoholCreate := "CREATE TABLE IF NOT EXISTS alcohol (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, alcohol_type_id INT NOT NULL, PRIMARY KEY (id), KEY alcohol_type_id (alcohol_type_id))"
	alcoholQuantityCreate := "CREATE TABLE IF NOT EXISTS alcohol_quantity (id INT NOT NULL AUTO_INCREMENT, alcohol_id INT NOT NULL, alcohol_quantity_type_id INT NOT NULL, price DECIMAL(10, 2) NOT NULL, PRIMARY KEY (id), KEY alcohol_id (alcohol_id), KEY alcohol_quantity_type_id (alcohol_quantity_type_id))"

	//Events
	venueCreate := "CREATE TABLE IF NOT EXISTS venue (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, location_id INT NOT NULL, PRIMARY KEY (id), KEY location_id (location_id))"
	venuePhotoCreate := "CREATE TABLE IF NOT EXISTS venue_photo (id INT NOT NULL AUTO_INCREMENT, venue_id INT NOT NULL, photo_url VARCHAR(255) NOT NULL, PRIMARY KEY (id), KEY venue_id (venue_id))"
	venueTimeblockCreate := "CREATE TABLE IF NOT EXISTS venue_timeblock (id INT NOT NULL AUTO_INCREMENT, venue_id INT NOT NULL, start_time DATETIME NOT NULL, end_time DATETIME NOT NULL, note VARCHAR(255), event_id INT, PRIMARY KEY (id), KEY venue_id (venue_id), KEY event_id (event_id))"

	eventTypeCreate := "CREATE TABLE IF NOT EXISTS event_type (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL UNIQUE, PRIMARY KEY (id))"
	venueEventTypeCreate := "CREATE TABLE IF NOT EXISTS venue_event_type (id INT NOT NULL AUTO_INCREMENT, venue_id INT NOT NULL, event_type_id INT NOT NULL, PRIMARY KEY (id), KEY venue_id (venue_id), KEY event_type_id (event_type_id))"
	venueEventTypeDefaultSettingsCreate := "CREATE TABLE IF NOT EXISTS venue_event_type_default_settings (id INT NOT NULL AUTO_INCREMENT, venue_event_type_id INT NOT NULL UNIQUE, hourly_rate DECIMAL(10, 2), minimum_booking_duration INT, flat_fee DECIMAL(10, 2), earliest_booking_time TIME NOT NULL, latest_booking_time TIME NOT NULL, PRIMARY KEY (id))"

	eventCreate := "CREATE TABLE IF NOT EXISTS event (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255), booking_id INT NOT NULL, venue_event_type_id INT NOT NULL, venue_timeblock_id INT NOT NULL, PRIMARY KEY (id), KEY booking_id (booking_id), KEY venue_event_type_id (venue_event_type_id), KEY venue_timeblock_id (venue_timeblock_id))"
	eventDetailsCreate := "CREATE TABLE IF NOT EXISTS event_details (id INT NOT NULL AUTO_INCREMENT, event_id INT NOT NULL UNIQUE, open_bar_requested BOOLEAN NOT NULL, alcohol_minimum DECIMAL(10, 2), guests INT, notes VARCHAR(255), PRIMARY KEY (id))"
	eventBookingCostCreate := "CREATE TABLE IF NOT EXISTS event_booking_cost (id INT NOT NULL AUTO_INCREMENT, event_id INT NOT NULL, booking_cost_item_id INT NOT NULL, PRIMARY KEY (id), KEY event_id (event_id), KEY booking_cost_item_id (booking_cost_item_id))"

	//file
	fileCreate := "CREATE TABLE IF NOT EXISTS file (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(255) NOT NULL, url VARCHAR(255) NOT NULL, PRIMARY KEY (id))"

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

	//Photo
	_, err = db.Exec(photoCreate)
	if err != nil {
		log.Fatalf("failed to create photo table: %v", err)
	}

	// Rentals

	_, err = db.Exec(rentalCreate)
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

	// Rental Photo
	_, err = db.Exec(rentalPhotoCreate)
	if err != nil {
		log.Fatalf("failed to create rental_photo table: %v", err)
	}

	// Boats
	_, err = db.Exec(boatCreate)
	if err != nil {
		log.Fatalf("failed to create boat table: %v", err)
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

	// Boat Photo
	_, err = db.Exec(boatPhotoCreate)
	if err != nil {
		log.Fatalf("failed to create boat_photo table: %v", err)
	}

	// Location
	_, err = db.Exec(locationCreate)
	if err != nil {
		log.Fatalf("failed to create location table: %v", err)
	}

	// User
	_, err = db.Exec(userCreate)
	if err != nil {
		log.Fatalf("failed to create user table: %v", err)
	}

	//Booking Status
	_, err = db.Exec(bookingStatusCreate)
	if err != nil {
		log.Fatalf("failed to create booking_status table: %v", err)
	}

	//Bookings
	_, err = db.Exec(bookingCreate)
	if err != nil {
		log.Fatalf("failed to create booking table: %v", err)
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
	_, err = db.Exec(rentalBookingCostCreate)
	if err != nil {
		log.Fatalf("failed to create rental_booking_cost table: %v", err)
	}

	//Boat Booking
	_, err = db.Exec(boatBookingCreate)
	if err != nil {
		log.Fatalf("failed to create boat_booking table: %v", err)
	}

	//Boat Booking Costs
	_, err = db.Exec(boatBookingCostCreate)
	if err != nil {
		log.Fatalf("failed to create boat_booking_cost table: %v", err)
	}

	//Booking Cost Type
	_, err = db.Exec(bookingCostTypeCreate)
	if err != nil {
		log.Fatalf("failed to create booking_cost_type table: %v", err)
	}

	//Booking Cost Items
	_, err = db.Exec(bookingCostItemCreate)
	if err != nil {
		log.Fatalf("failed to create booking_cost_item table: %v", err)
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

	//Refund Status
	_, err = db.Exec(refundStatusCreate)
	if err != nil {
		log.Fatalf("failed to create refund_status table: %v", err)
	}

	//Refund Requests
	_, err = db.Exec(refundRequestCreate)
	if err != nil {
		log.Fatalf("failed to create refund_request table: %v", err)
	}

	//Alcohol Type
	_, err = db.Exec(alcoholTypeCreate)
	if err != nil {
		log.Fatalf("failed to create alcohol_type table: %v", err)
	}

	//Alcohol Quantity Type
	_, err = db.Exec(alcoholQuantityTypeCreate)
	if err != nil {
		log.Fatalf("failed to create alcohol_quantity_type table: %v", err)
	}

	//Alcohol
	_, err = db.Exec(alcoholCreate)
	if err != nil {
		log.Fatalf("failed to create alcohol table: %v", err)
	}

	//Alcohol Quantity
	_, err = db.Exec(alcoholQuantityCreate)
	if err != nil {
		log.Fatalf("failed to create alcohol_quantity table: %v", err)
	}

	//Alcohol Order
	_, err = db.Exec(alcoholOrderCreate)
	if err != nil {
		log.Fatalf("failed to create alcohol_order table: %v", err)
	}

	//Alcohol Order Item
	_, err = db.Exec(alcoholOrderItemCreate)
	if err != nil {
		log.Fatalf("failed to create alcohol_order_item table: %v", err)
	}

	//Alcohol Order Booking Cost
	_, err = db.Exec(alcoholOrderBookingCost)
	if err != nil {
		log.Fatalf("failed to create alcohol_order_booking_cost table: %v", err)
	}

	//Events
	_, err = db.Exec(venueCreate)
	if err != nil {
		log.Fatalf("failed to create venue table: %v", err)
	}

	//Venue Photo
	_, err = db.Exec(venuePhotoCreate)
	if err != nil {
		log.Fatalf("failed to create venue_photo table: %v", err)
	}

	//Venue Timeblock
	_, err = db.Exec(venueTimeblockCreate)
	if err != nil {
		log.Fatalf("failed to create venue_timeblock table: %v", err)
	}

	//Event Type
	_, err = db.Exec(eventTypeCreate)
	if err != nil {
		log.Fatalf("failed to create event_type table: %v", err)

	}

	//Venue Event Type
	_, err = db.Exec(venueEventTypeCreate)
	if err != nil {
		log.Fatalf("failed to create venue_event_type table: %v", err)
	}

	//Venue Event Type Default Settings
	_, err = db.Exec(venueEventTypeDefaultSettingsCreate)
	if err != nil {
		log.Fatalf("failed to create venue_event_type_default_settings table: %v", err)
	}

	//Event
	_, err = db.Exec(eventCreate)
	if err != nil {
		log.Fatalf("failed to create event table: %v", err)
	}

	//Event Details
	_, err = db.Exec(eventDetailsCreate)
	if err != nil {
		log.Fatalf("failed to create event_details table: %v", err)
	}

	//Event Booking Cost
	_, err = db.Exec(eventBookingCostCreate)
	if err != nil {
		log.Fatalf("failed to create event_booking_cost table: %v", err)
	}

	//File
	_, err = db.Exec(fileCreate)
	if err != nil {
		log.Fatalf("failed to create file table: %v", err)
	}

	//Booking File
	_, err = db.Exec(bookingFileCreate)
	if err != nil {
		log.Fatalf("failed to create booking_file table: %v", err)
	}

	//Boat Variable Settings

	_, err = db.Exec(boatVariableSettingsCreate)
	if err != nil {
		log.Fatalf("failed to create boat_variable_settings table: %v", err)
	}

	//Rental Bedrooms
	_, err = db.Exec(bedTypeCreate)
	if err != nil {
		log.Fatalf("failed to create bed_type table: %v", err)
	}

	_, err = db.Exec(rentalBedroomCreate)
	if err != nil {
		log.Fatalf("failed to create rental_bedroom table: %v", err)
	}

	_, err = db.Exec(rentalBedroomBedCreate)
	if err != nil {
		log.Fatalf("failed to create rental_bedroom_bed table: %v", err)
	}

	//Rental Bathrooms
	_, err = db.Exec(rentallBathroomCreate)

	if err != nil {
		log.Fatalf("failed to create rental_bathroom table: %v", err)
	}

	//Amenities
	_, err = db.Exec(amenityTypeCreate)
	if err != nil {
		log.Fatalf("failed to create amenity_type table: %v", err)
	}

	_, err = db.Exec(amenityCreate)
	if err != nil {
		log.Fatalf("failed to create amenity table: %v", err)
	}

	_, err = db.Exec(rentalAmenityCreate)
	if err != nil {
		log.Fatalf("failed to create rental_amenity table: %v", err)
	}

	// Close the connection

	defer db.Close()
}
