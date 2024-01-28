package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type BoatVariableSettings struct {
	ID                     int
	BoatID                 int
	StartDate              time.Time
	EndDate                time.Time
	DailyCost              float64
	MinimumBookingDuration int
}

func GetVariableSettingsForBoatId(boatId int, db *sql.DB) ([]BoatVariableSettings, error) {
	var boatVariableSettings []BoatVariableSettings

	rows, err := db.Query("SELECT id, boat_id, start_date, end_date, daily_cost, minimum_booking_duration FROM boat_variable_settings WHERE boat_id = ?", boatId)
	if err != nil {
		return boatVariableSettings, err
	}
	defer rows.Close()

	for rows.Next() {
		var boatVariableSetting BoatVariableSettings

		var startDateString string
		var endDateString string
		err := rows.Scan(&boatVariableSetting.ID, &boatVariableSetting.BoatID, &startDateString, &endDateString, &boatVariableSetting.DailyCost, &boatVariableSetting.MinimumBookingDuration)
		if err != nil {
			return boatVariableSettings, err
		}

		boatVariableSetting.StartDate, err = time.Parse("2006-01-02", startDateString)
		if err != nil {
			return boatVariableSettings, err
		}

		boatVariableSetting.EndDate, err = time.Parse("2006-01-02", endDateString)
		if err != nil {
			return boatVariableSettings, err
		}

		boatVariableSettings = append(boatVariableSettings, boatVariableSetting)
	}
	err = rows.Err()
	if err != nil {
		return boatVariableSettings, err
	}

	return boatVariableSettings, nil

}
func AttemptToCreateVariableSettingsForBoatId(boatId int, startDate time.Time, endDate time.Time, dailyCost float64, minimumBookingDuration int, db *sql.DB) (int64, error) {
	//ensure there is no overlap
	boatVariableSettings, err := GetVariableSettingsForBoatId(boatId, db)
	if err != nil {
		return 0, err
	}

	for _, boatVariableSetting := range boatVariableSettings {
		if startDate.Before(boatVariableSetting.EndDate) && endDate.After(boatVariableSetting.StartDate) {
			return 0, nil
		}
	}

	result, err := db.Exec("INSERT INTO boat_variable_settings (boat_id, start_date, end_date, daily_cost, minimum_booking_duration) VALUES (?, ?, ?, ?, ?)", boatId, startDate, endDate, dailyCost, minimumBookingDuration)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetVariableSettingsForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	boatId := vars["id"]

	boatIdInt, err := strconv.Atoi(boatId)
	boatVariableSettings, err := GetVariableSettingsForBoatId(boatIdInt, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(boatVariableSettings)
}

func GetVariableSettingsForBoatIdAndDateRange(boatId string, startDate time.Time, endDate time.Time, db *sql.DB) ([]BoatVariableSettings, error) {
	rows, err := db.Query("SELECT id, boat_id, start_date, end_date, daily_cost, minimum_booking_duration FROM boat_variable_settings WHERE boat_id = ? AND start_date <= ? AND end_date >= ?", boatId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boatVariableSettings []BoatVariableSettings

	for rows.Next() {
		var boatVariableSetting BoatVariableSettings

		var startDateString string
		var endDateString string
		err := rows.Scan(&boatVariableSetting.ID, &boatVariableSetting.BoatID, &startDateString, &endDateString, &boatVariableSetting.DailyCost, &boatVariableSetting.MinimumBookingDuration)
		if err != nil {
			return nil, err
		}

		boatVariableSetting.StartDate, err = time.Parse("2006-01-02", startDateString)
		if err != nil {
			return nil, err
		}

		boatVariableSetting.EndDate, err = time.Parse("2006-01-02", endDateString)
		if err != nil {
			return nil, err
		}

		boatVariableSettings = append(boatVariableSettings, boatVariableSetting)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return boatVariableSettings, nil

}
func CreateVariableSettingsForBoat(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	boatId := vars["id"]

	boatIdInt, err := strconv.Atoi(boatId)

	var boatVariableSettings BoatVariableSettings
	err = json.NewDecoder(r.Body).Decode(&boatVariableSettings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = AttemptToCreateVariableSettingsForBoatId(boatIdInt, boatVariableSettings.StartDate, boatVariableSettings.EndDate, boatVariableSettings.DailyCost, boatVariableSettings.MinimumBookingDuration, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
