package api

import (
	"booking-api/config"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BoatPhoto struct {
	ID       int
	BoatID   int
	PhotoURL string
}

func GetBoatThumbnailByBoatID(boatID int, db *sql.DB) (string, error) {
	boatIDString := strconv.Itoa(boatID)
	rows, err := db.Query("SELECT photo_url FROM boat_photo WHERE boat_id = ? LIMIT 1", boatIDString)
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

func GetBoatPhotosByID(boatID string, db *sql.DB) ([]BoatPhoto, error) {
	rows, err := db.Query("SELECT id, boat_id, photo_url FROM boat_photo WHERE boat_id = ?", boatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	boatPhotos := []BoatPhoto{}

	for rows.Next() {
		var boatPhoto BoatPhoto
		if err := rows.Scan(&boatPhoto.ID, &boatPhoto.BoatID, &boatPhoto.PhotoURL); err != nil {
			return nil, err
		}
		boatPhotos = append(boatPhotos, boatPhoto)
	}

	return boatPhotos, nil
}

func GetBoatPhotos(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	boatID := vars["id"]

	boatPhotos, err := GetBoatPhotosByID(boatID, db)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting boat photos", http.StatusInternalServerError)
		return
	}

	// Load connection string from .env file
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	for i := range boatPhotos {
		boatPhotos[i].PhotoURL = env.OBJECT_STORAGE_URL + "/" + boatPhotos[i].PhotoURL
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boatPhotos)
}

func CreateBoatPhoto(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	boatID := mux.Vars(r)["id"]

	newFilePath := "boat_photos/" + boatID

	uploadedFilePath, err := UploadHandler(w, r, newFilePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error uploading photo", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO boat_photo (boat_id, photo_url) VALUES (?, ?)", boatID, uploadedFilePath)
	if err != nil {

		log.Println(err)
		http.Error(w, "Error inserting photo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func DeleteBoatPhoto(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	boatID := mux.Vars(r)["id"]
	photoID := mux.Vars(r)["photoID"]

	//delete the photo from object storage
	var photoURL string

	err := db.QueryRow("SELECT photo_url FROM boat_photo WHERE id = ? AND boat_id = ?", photoID, boatID).Scan(&photoURL)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting photo", http.StatusInternalServerError)
		return
	}

	err = DeleteHandler(w, r, photoURL)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error deleting photo", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM boat_photo WHERE id = ? AND boat_id = ?", photoID, boatID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error deleting photo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
