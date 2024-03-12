package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBoatThumbnail(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	idStr, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("failed to convert id to int: %v", err)
	}

	thumbnail, err := GetBoatThumbnailByBoatID(idStr, db)

	if err != nil {
		log.Fatalf("failed to get thumbnail: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thumbnail)

}
