package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/objectStorage"

	"gorm.io/gorm"
)

type PhotoRepositoryImplementation struct {
	StorageClient *objectStorage.S3Client
	DB            *gorm.DB
}

func NewPhotoRepositoryImplementation(storageClient *objectStorage.S3Client, db *gorm.DB) *PhotoRepositoryImplementation {
	return &PhotoRepositoryImplementation{StorageClient: storageClient, DB: db}
}

// FindAll retrieves all photos from the database
func (r *PhotoRepositoryImplementation) FindAll() []response.PhotoResponse {

	var photos []models.Photo
	result := r.DB.Model(&photos).Find(&photos)

	if result.Error != nil {
		return []response.PhotoResponse{}
	}

	var photoResponses []response.PhotoResponse
	for _, photo := range photos {
		photoResponses = append(photoResponses, response.PhotoResponse{
			ID:  photo.ID,
			URL: photo.URL,
		})
	}

	return photoResponses

}
