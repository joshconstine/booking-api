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
	FileID                 int
}

func GetSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query the database for the default setting of the id.
	query := "SELECT * FROM rental_unit_default_settings WHERE rental_id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	// Create a slice of RentalUnitDefaultSettings to hold the data.
	rentalUnitDefualtSettings := []RentalUnitDefaultSettings{}

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var rentalUnitDefualtSetting RentalUnitDefaultSettings
		err := rows.Scan(
			&rentalUnitDefualtSetting.ID,
			&rentalUnitDefualtSetting.RentalID,
			&rentalUnitDefualtSetting.NightlyCost,
			&rentalUnitDefualtSetting.MinimumBookingDuration,
			&rentalUnitDefualtSetting.AllowsPets,
			&rentalUnitDefualtSetting.CleaningFee,
			&rentalUnitDefualtSetting.CheckInTime,
			&rentalUnitDefualtSetting.CheckOutTime,
			&rentalUnitDefualtSetting.FileID,
		)
		if err != nil {
			log.Fatalf("failed to scan: %v", err)
		}

		// Append the struct to the slice.
		rentalUnitDefualtSettings = append(rentalUnitDefualtSettings, rentalUnitDefualtSetting)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalUnitDefualtSettings)

}
