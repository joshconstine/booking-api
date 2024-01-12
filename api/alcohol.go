package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Alcohol struct {
	ID            int
	Name          string
	AlcoholTypeId int
}

type OrderableAlcohol struct {
	ID                      int
	Name                    string
	AlcoholType             string
	AlcoholTypeID           int
	Price                   float32
	AlcoholQuantityTypeID   int
	AlcoholQuantityTypeName string
}

type AlcoholType struct {
	ID   int
	Name string
}

type AlcoholQuantityType struct {
	ID   int
	Name string
}

func GetAlcohol(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all the alcohol from the database.

	rows, err := db.Query("SELECT * FROM alcohol")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var alcohols []Alcohol

	for rows.Next() {
		var alcohol Alcohol
		if err := rows.Scan(&alcohol.ID, &alcohol.Name, &alcohol.AlcoholTypeId); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		alcohols = append(alcohols, alcohol)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alcohols)

}

func GetOrderableAlcohol(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all the alcohol from the database.

	rows, err := db.Query("SELECT alcohol.id, alcohol.name, alcohol_type.name, alcohol.alcohol_type_id, alcohol_quantity.price, alcohol_quantity_type.id, alcohol_quantity_type.name FROM alcohol_quantity JOIN alcohol ON alcohol_quantity.alcohol_id = alcohol.id JOIN alcohol_type ON alcohol.alcohol_type_id = alcohol_type.id JOIN alcohol_quantity_type ON alcohol_quantity.alcohol_quantity_type_id = alcohol_quantity_type.id")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var alcohols []OrderableAlcohol

	for rows.Next() {
		var alcohol OrderableAlcohol
		if err := rows.Scan(&alcohol.ID, &alcohol.Name, &alcohol.AlcoholType, &alcohol.AlcoholTypeID, &alcohol.Price, &alcohol.AlcoholQuantityTypeID, &alcohol.AlcoholQuantityTypeName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		alcohols = append(alcohols, alcohol)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alcohols)

}

func GetAlcoholTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all the alcohol from the database.

	rows, err := db.Query("SELECT * FROM alcohol_type")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var alcoholTypes []AlcoholType

	for rows.Next() {
		var alcoholType AlcoholType
		if err := rows.Scan(&alcoholType.ID, &alcoholType.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		alcoholTypes = append(alcoholTypes, alcoholType)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alcoholTypes)

}

func GetAlcoholQuantityTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all the alcohol from the database.

	rows, err := db.Query("SELECT * FROM alcohol_quantity_type")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var alcoholQuantityTypes []AlcoholQuantityType

	for rows.Next() {
		var alcoholQuantityType AlcoholQuantityType
		if err := rows.Scan(&alcoholQuantityType.ID, &alcoholQuantityType.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		alcoholQuantityTypes = append(alcoholQuantityTypes, alcoholQuantityType)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alcoholQuantityTypes)

}
