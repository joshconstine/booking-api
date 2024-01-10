package main

import (
    "booking-api/api"
    "database/sql"
    "log"
    "os"
    "fmt"
    "strconv"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
     _ "github.com/go-sql-driver/mysql"
)

func main() {

    port := 8080
	// Convert the integer port to a string.
	portStr := strconv.Itoa(port)
	
	

	r := mux.NewRouter()
	
    // Use the functions from the 'api' package to define routes.
    
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
    
    api.InitRoutes(r, db)

    fmt.Printf("Server is listening on port %d...\n", port)

    http.ListenAndServe(":"+portStr, r)


}
