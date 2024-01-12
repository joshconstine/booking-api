package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RentalUnitDefaultSettings struct {
	ID                     int
	RentalID               int
	NightlyCost            float64
	MinimumBookingDuration int
	AllowsPets             bool
	CleaningFee            float64
	CheckInTime            string
	CheckOutTime           string
}

func GetSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	rows, err := db.Query("SELECT rental_unit_default_settings.id, rental_unit_default_settings.rental_id, rental_unit_default_settings.nightly_cost, rental_unit_default_settings.minimum_booking_duration, rental_unit_default_settings.allows_pets, rental_unit_default_settings.cleaning_fee, rental_unit_default_settings.check_in_time, rental_unit_default_settings.check_out_time FROM rental_unit_default_settings JOIN rentals ON rental_unit_default_settings.rental_id = rentals.id WHERE rentals.id = ?", id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of rentals to hold the data.
	var settings []RentalUnitDefaultSettings

	// Loop through the data and insert into the rentals slice.
	for rows.Next() {
		var setting RentalUnitDefaultSettings
		if err := rows.Scan(&setting.ID, &setting.RentalID, &setting.NightlyCost, &setting.MinimumBookingDuration, &setting.AllowsPets, &setting.CleaningFee, &setting.CheckInTime, &setting.CheckOutTime); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		settings = append(settings, setting)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)

}
