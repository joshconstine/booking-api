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

type RentalPhoto struct {
	ID       int
	RentalID int
	PhotoURL string
}

func GetRentalThumbnailByRentalID(rentalID int, db *sql.DB) (string, error) {
	rentalIDString := strconv.Itoa(rentalID)
	rows, err := db.Query("SELECT photo_url FROM rental_photo WHERE rental_id = ? LIMIT 1", rentalIDString)
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

func getRentalPhotos(rentalID string, db *sql.DB) ([]RentalPhoto, error) {
	rows, err := db.Query("SELECT id, rental_id, photo_url FROM rental_photo WHERE rental_id = ?", rentalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rentalPhotos := []RentalPhoto{}

	for rows.Next() {
		var rentalPhoto RentalPhoto
		if err := rows.Scan(&rentalPhoto.ID, &rentalPhoto.RentalID, &rentalPhoto.PhotoURL); err != nil {
			return nil, err
		}
		// Load connection string from .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to load env", err)
		}

		rentalPhoto.PhotoURL = os.Getenv("OBJECT_STORAGE_URL") + "/" + rentalPhoto.PhotoURL

		rentalPhotos = append(rentalPhotos, rentalPhoto)
	}

	return rentalPhotos, nil
}

func GetRentalPhotos(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	rentalID := vars["id"]
	rentalPhotos, err := getRentalPhotos(rentalID, db)
	if err != nil {
		log.Fatalf("failed to get rental photos: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentalPhotos)

}
