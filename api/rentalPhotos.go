package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type RentalPhoto struct {
	ID       int
	RentalID int
	PhotoURL string
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
