package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type BedType struct {
	ID   int
	Name string
}

func GetBedTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//
	// Get all bed types.

	rows, err := db.Query("SELECT id, name FROM bed_type")
	if err != nil {
		log.Fatalf("failed to get bed types: %v", err)
	}

	var bedTypes []BedType
	for rows.Next() {
		var bedType BedType
		if err := rows.Scan(&bedType.ID, &bedType.Name); err != nil {
			log.Fatalf("failed to scan bed type: %v", err)
		}
		bedTypes = append(bedTypes, bedType)
	}

	if err := json.NewEncoder(w).Encode(bedTypes); err != nil {
		log.Fatalf("failed to encode bed types: %v", err)
	}
}
