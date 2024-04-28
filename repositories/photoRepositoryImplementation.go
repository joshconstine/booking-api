package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/pkg/objectStorage"
	"mime/multipart"

	"gorm.io/gorm"
)

type PhotoRepositoryImplementation struct {
	StorageClient *objectStorage.S3Client
	DB            *gorm.DB
}

func NewPhotoRepositoryImplementation(storageClient *objectStorage.S3Client, db *gorm.DB) PhotoRepository {
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

// AddPhoto adds a photo to the database
func (r *PhotoRepositoryImplementation) AddPhoto(filePath string, photo *multipart.File, fileExt string) response.PhotoResponse {

	returnedUrl, err := r.StorageClient.UploadFile(photo, filePath, fileExt)

	if err != nil {
		return response.PhotoResponse{}
	}

	var photoModel models.Photo
	photoModel.URL = returnedUrl

	result := r.DB.Create(&photoModel)

	if result.Error != nil {
		return response.PhotoResponse{}
	}

	return response.PhotoResponse{
		ID:  photoModel.ID,
		URL: photoModel.URL,
	}

}
