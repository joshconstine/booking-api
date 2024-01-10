package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

// InitRoutes initializes the routes for the API.
func InitRoutes(r *mux.Router, db *sql.DB) {
	// Define the routes.



    r.HandleFunc("/rentals",  func(w http.ResponseWriter, r *http.Request) {
        GetRentals(w, r, db )
    }).Methods("GET") 

	r.HandleFunc("/rentals/{id}",  func(w http.ResponseWriter, r *http.Request) {
		GetRental(w, r, db )
}).Methods("GET")


}