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
	env, err := config.LoadConfig("../")
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

func CreateVenuePhoto(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	venueID := mux.Vars(r)["id"]

	newFilePath := "venue_photos/" + venueID

	uploadedFilePath, err := UploadHandler(w, r, newFilePath)
	if err != nil {
		log.Fatalf("failed to upload venue photo: %v", err)
		http.Error(w, "Failed to upload venue photo", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO venue_photo (venue_id, photo_url) VALUES (?, ?)", venueID, uploadedFilePath)
	if err != nil {
		log.Fatalf("failed to insert venue photo into database: %v", err)
		http.Error(w, "Failed to insert venue photo into database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func DeleteVenuePhoto(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	photoID := vars["photoID"]

	var photoURL string
	err := db.QueryRow("SELECT photo_url FROM venue_photo WHERE id = ?", photoID).Scan(&photoURL)
	if err != nil {
		log.Fatalf("failed to get venue photo: %v", err)
		http.Error(w, "Failed to get venue photo", http.StatusInternalServerError)
		return
	}

	err = DeleteHandler(w, r, photoURL)
	if err != nil {
		log.Fatalf("failed to delete venue photo: %v", err)
		http.Error(w, "Failed to delete venue photo", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM venue_photo WHERE id = ?", photoID)
	if err != nil {
		log.Fatalf("failed to delete venue photo: %v", err)
		http.Error(w, "Failed to delete venue photo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
