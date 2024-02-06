package api

import (
	"booking-api/config"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type VenuePhoto struct {
	ID       int
	VenueID  int
	PhotoURL string
}

func GetVenueThumbnailByVenueID(venueID int, db *sql.DB) (string, error) {
	venueIDString := strconv.Itoa(venueID)
	rows, err := db.Query("SELECT photo_url FROM venue_photo WHERE venue_id = ? LIMIT 1", venueIDString)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var photoURL string

	for rows.Next() {
		if err := rows.Scan(&photoURL); err != nil {
			return "", err
		}
	}

	// Load connection string from .env file
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	photoURL = env.OBJECT_STORAGE_URL + "/" + photoURL

	return photoURL, nil
}

func getVenuePhotos(venueID string, db *sql.DB) ([]VenuePhoto, error) {
	rows, err := db.Query("SELECT id, venue_id, photo_url FROM venue_photo WHERE venue_id = ?", venueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	venuePhotos := []VenuePhoto{}

	for rows.Next() {
		var venuePhoto VenuePhoto
		if err := rows.Scan(&venuePhoto.ID, &venuePhoto.VenueID, &venuePhoto.PhotoURL); err != nil {
			return nil, err
		}
		// Load connection string from .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to load env", err)
		}

		venuePhoto.PhotoURL = os.Getenv("OBJECT_STORAGE_URL") + "/" + venuePhoto.PhotoURL

		venuePhotos = append(venuePhotos, venuePhoto)
	}

	return venuePhotos, nil
}

func GetVenuePhotos(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	venueID := vars["id"]
	venuePhotos, err := getVenuePhotos(venueID, db)
	if err != nil {
		log.Fatalf("failed to get venue photos: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venuePhotos)

}
