

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
	dropTables := "DROP TABLE IF EXISTS rentals, rental_timeblock, location, rental_unit_default_settings, rental_unit_variable_settings, rental_photo, photo, user, booking_status ,boat_booking,boat_booking_costs, booking, booking_details, rental_booking, rental_booking_cost, booking_payment, booking_cost_type, booking_cost_item, boats, boat_timeblock, boat_photo, refund_status, refund_request, rental_status, boat_status, boat_default_settings, payment_method, alcohol, alcohol_type, alcohol_quantity_type, alcohol_quantity, alcohol_order, alcohol_order_booking_cost, alcohol_order_item"

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


	_, err = db.Exec(dropTables)
	if err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}
	

	log.Println("dropped tables")


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
