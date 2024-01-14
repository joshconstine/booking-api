package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type RentalUnitVariableSettings struct {
	ID                     int
	RentalID               int
	StartDate              time.Time
	EndDate                time.Time
	MinimumBookingDuration int
	NightlyCost            float64
	CleaningFee            float64
	EventRequired          bool
}

func GetVariableSettingsForRentalId(rentalId string, db *sql.DB) ([]RentalUnitVariableSettings, error) {
	rows, err := db.Query("SELECT id, rental_id, nightly_cost, minimum_booking_duration, cleaning_fee, event_required, start_date, end_date FROM rental_unit_variable_settings WHERE rental_id = ?", rentalId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentalUnitVariableSettings []RentalUnitVariableSettings

	for rows.Next() {
		var startDateStr, endDateStr string // Change types to string

		var rentalUnitVariableSetting RentalUnitVariableSettings

		err := rows.Scan(
			&rentalUnitVariableSetting.ID,
			&rentalUnitVariableSetting.RentalID,
			&rentalUnitVariableSetting.NightlyCost,
			&rentalUnitVariableSetting.MinimumBookingDuration,
			&rentalUnitVariableSetting.CleaningFee,
			&rentalUnitVariableSetting.EventRequired,
			&startDateStr,
			&endDateStr,
		)
		if err != nil {
			return nil, err
		}

		// Parse date strings to time.Time
		rentalUnitVariableSetting.StartDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return nil, err
		}

		rentalUnitVariableSetting.EndDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return nil, err
		}

		rentalUnitVariableSettings = append(rentalUnitVariableSettings, rentalUnitVariableSetting)
	}

	return rentalUnitVariableSettings, nil
}

func GetVariableSettingsForRentalIdAndDate(rentalId string, date time.Time, db *sql.DB) (RentalUnitVariableSettings, error) {

	rows, err := db.Query("SELECT * FROM rental_unit_variable_settings WHERE rental_id = ? AND start_date <= ? AND end_date >= ?", rentalId, date, date)

	if err != nil {
		return RentalUnitVariableSettings{}, err
	}

	defer rows.Close()

	var rentalUnitVariableSettings RentalUnitVariableSettings

	if rows.Next() {
		err := rows.Scan(&rentalUnitVariableSettings.ID, &rentalUnitVariableSettings.RentalID, &rentalUnitVariableSettings.NightlyCost, &rentalUnitVariableSettings.MinimumBookingDuration, &rentalUnitVariableSettings.CleaningFee, &rentalUnitVariableSettings.EventRequired, &rentalUnitVariableSettings.StartDate, &rentalUnitVariableSettings.EndDate)
		if err != nil {
			return RentalUnitVariableSettings{}, err
		}
	}

	return rentalUnitVariableSettings, nil
}

func GetVariableSettingsForRentalIdAndDateRange(rentalId string, startDate time.Time, endDate time.Time, db *sql.DB) ([]RentalUnitVariableSettings, error) {

	rows, err := db.Query("SELECT * FROM rental_unit_variable_settings WHERE rental_id = ? AND start_date <= ? AND end_date >= ?", rentalId, startDate, endDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rentalUnitVariableSettings []RentalUnitVariableSettings

	for rows.Next() {
		var rentalUnitVariableSetting RentalUnitVariableSettings
		err := rows.Scan(&rentalUnitVariableSetting.ID, &rentalUnitVariableSetting.RentalID, &rentalUnitVariableSetting.NightlyCost, &rentalUnitVariableSetting.MinimumBookingDuration, &rentalUnitVariableSetting.CleaningFee, &rentalUnitVariableSetting.EventRequired, &rentalUnitVariableSetting.StartDate, &rentalUnitVariableSetting.EndDate)
		if err != nil {
			return nil, err
		}
		rentalUnitVariableSettings = append(rentalUnitVariableSettings, rentalUnitVariableSetting)
	}

	return rentalUnitVariableSettings, nil
}

func UpdateVariableSettingsForRentalId(rentalId string, nightlyCost float64, minimumBookingDuration int, cleaningFee float64, eventRequired bool, StartDate time.Time, EndDate time.Time, db *sql.DB) error {

	query := "UPDATE rental_unit_variable_settings SET nightly_cost = ?, minimum_booking_duration = ?, cleaning_fee = ?, event_required = ?, start_date = ?, end_date = ? WHERE rental_id = ?"

	_, err := db.Exec(query, nightlyCost, minimumBookingDuration, cleaningFee, eventRequired, StartDate, EndDate, rentalId)

	if err != nil {
		return err
	}

	return nil

}

func CreateVariableSettingsForRentalId(rentalId string, nightlyCost float64, minimumBookingDuration int, cleaningFee float64, eventRequired bool, StartDate time.Time, EndDate time.Time, db *sql.DB) error {

	query := "INSERT INTO rental_unit_variable_settings (rental_id, nightly_cost, minimum_booking_duration, cleaning_fee, event_required, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query, rentalId, nightlyCost, minimumBookingDuration, cleaningFee, eventRequired, StartDate, EndDate)

	if err != nil {
		return err
	}

	return nil

}

func DeleteVariableSettingsForRentalId(rentalId string, db *sql.DB) error {

	query := "DELETE FROM rental_unit_variable_settings WHERE rental_id = ?"

	_, err := db.Exec(query, rentalId)

	if err != nil {
		return err
	}

	return nil

}

// API functions

func GetVariableSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	rentalID := vars["id"]

	rentalUnitVariableSettings, err := GetVariableSettingsForRentalId(rentalID, db)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalUnitVariableSettings)

}

func UpdateVariableSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	rentalID := vars["id"]

	var rentalUnitVariableSettings RentalUnitVariableSettings
	err := json.NewDecoder(r.Body).Decode(&rentalUnitVariableSettings)
	if err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	err = UpdateVariableSettingsForRentalId(rentalID, rentalUnitVariableSettings.NightlyCost, rentalUnitVariableSettings.MinimumBookingDuration, rentalUnitVariableSettings.CleaningFee, rentalUnitVariableSettings.EventRequired, rentalUnitVariableSettings.StartDate, rentalUnitVariableSettings.EndDate, db)
	if err != nil {
		log.Fatalf("failed to update: %v", err)
	}

}

func CreateVariableSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	rentalID := vars["id"]

	var rentalUnitVariableSettings RentalUnitVariableSettings
	err := json.NewDecoder(r.Body).Decode(&rentalUnitVariableSettings)
	if err != nil {
		log.Fatalf("failed to decode: %v", err)
	}

	err = CreateVariableSettingsForRentalId(rentalID, rentalUnitVariableSettings.NightlyCost, rentalUnitVariableSettings.MinimumBookingDuration, rentalUnitVariableSettings.CleaningFee, rentalUnitVariableSettings.EventRequired, rentalUnitVariableSettings.StartDate, rentalUnitVariableSettings.EndDate, db)
	if err != nil {
		log.Fatalf("failed to create: %v", err)
	}

}

func DeleteVariableSettingsForRental(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	rentalID := vars["id"]

	err := DeleteVariableSettingsForRentalId(rentalID, db)
	if err != nil {
		log.Fatalf("failed to delete: %v", err)
	}

}
