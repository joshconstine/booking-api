package api

import (
	"booking-api/config"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
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

func CreateRentalPhoto(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rentalID := mux.Vars(r)["id"]
	var bucket, key string
	var timeout time.Duration

	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	bucket = env.OBJECT_STORAGE_BUCKET

	// key = "rental_photos/" + rentalID + "/" + uuid.New().String()

	key = "rental_photos/" + rentalID

	// flag.StringVar(&bucket, "b", "", "Bucket name.")
	// flag.StringVar(&key, "k", "", "Object key name.")
	flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
	flag.Parse()

	ACCESS_KEY := env.OBJECT_STORAGE_ACCESS_KEY
	SECRET_KEY := env.OBJECT_STORAGE_SECRET

	//log the access key
	log.Printf("access key: %v", ACCESS_KEY)

	//log the secret key
	log.Printf("secret key: %v", SECRET_KEY)

	config := &aws.Config{
		Region:      aws.String("us-ord-1"),
		Endpoint:    aws.String("https://us-ord-1.linodeobjects.com"), // Linode Object Storage endpoint
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
	}

	//log the config

	log.Printf("config: %v", config)

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	sess := session.Must(session.NewSession(config))

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := s3.New(sess)

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	if cancelFn != nil {
		defer cancelFn()
	}
	err = r.ParseMultipartForm(10 * 1024 * 1024) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// Generate a unique filename using a UUID
	fileExt := filepath.Ext(header.Filename)
	newFilename := uuid.New().String() + fileExt

	// Create a new file in the "public/static" directory with the unique filename
	newFilePath := filepath.Join(key, newFilename)

	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(newFilePath),
		Body:        file,
		ContentType: aws.String(fileExt),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)

	newPhotoLocation := key
	//log the photo location
	log.Printf("photo location: %v", newPhotoLocation)

	//insert the photo location into the database

	_, err = db.Exec("INSERT INTO rental_photo (rental_id, photo_url) VALUES (?, ?)", rentalID, newPhotoLocation)
	if err != nil {
		log.Fatalf("failed to insert rental photo: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
