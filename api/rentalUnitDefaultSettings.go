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

func GetDefaultSettingsForRentalId(id string, db *sql.DB) (RentalUnitDefaultSettings, error) {
	// Query the database for the default setting of the id.
	query := "SELECT * FROM rental_unit_default_settings WHERE rental_id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return RentalUnitDefaultSettings{}, err
	}
	defer rows.Close()

	// Create a single instance of RentalUnitDefaultSettings.
	var rentalUnitDefualtSetting RentalUnitDefaultSettings

	// Check if there is at least one row.
	if rows.Next() {
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
			return RentalUnitDefaultSettings{}, err
		}
	}

	return rentalUnitDefualtSetting, nil
}

func GetDefaultSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	rentalUnitDefualtSetting, err := GetDefaultSettingsForRentalId(id, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalUnitDefualtSetting)
}
func UpdateDefaultSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Create a single instance of RentalUnitDefaultSettings.
	var rentalUnitDefualtSetting RentalUnitDefaultSettings

	// Decode the JSON data sent by the client.
	err := json.NewDecoder(r.Body).Decode(&rentalUnitDefualtSetting)
	if err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	// Update the database only for values that where passed in teh json.

	query := "UPDATE rental_unit_default_settings SET nightly_cost = ?, minimum_booking_duration = ?, allows_pets = ?, cleaning_fee = ?, check_in_time = ?, check_out_time = ?, file_id = ? WHERE rental_id = ?"
	_, err = db.Exec(query,
		rentalUnitDefualtSetting.NightlyCost,
		rentalUnitDefualtSetting.MinimumBookingDuration,
		rentalUnitDefualtSetting.AllowsPets,
		rentalUnitDefualtSetting.CleaningFee,
		rentalUnitDefualtSetting.CheckInTime,
		rentalUnitDefualtSetting.CheckOutTime,
		rentalUnitDefualtSetting.FileID,
		id,
	)
	if err != nil {
		log.Fatalf("failed to update: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalUnitDefualtSetting)
}
