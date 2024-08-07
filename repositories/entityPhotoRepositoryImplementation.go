package repositories

import (
	"booking-api/config"
	"booking-api/data/response"
	"booking-api/models"
	"fmt"

	"gorm.io/gorm"
)

type EntityPhotoRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityPhotoRepositoryImplementation(db *gorm.DB) EntityPhotoRepository {
	return &EntityPhotoRepositoryImplementation{Db: db}
}

func (e *EntityPhotoRepositoryImplementation) FindAllForEntity(entity string, entityID uint) []response.PhotoResponse {
	var entityPhotos []models.EntityPhoto
	result := e.Db.Model(&models.EntityPhoto{}).Where("entity_id = ? AND entity_type = ?", entityID, entity).Find(&entityPhotos)
	if result.Error != nil {
		return []response.PhotoResponse{}
	}

	var photos []response.PhotoResponse
	var ph response.PhotoResponse
	var photoModel models.Photo
	for _, entityPhoto := range entityPhotos {

		photo := e.Db.Model(&models.Photo{}).Where("id = ?", entityPhoto.PhotoID).First(&photoModel)

		if photo.Error == nil {
			ph.ID = entityPhoto.PhotoID
			ph.URL = entityPhoto.Photo.URL
		}

		ph.ID = photoModel.ID
		ph.URL = photoModel.URL

		// load config
		env, err := config.LoadConfig(".")
		if err != nil {
			fmt.Printf("error: %v", err)
		}

		base := env.OBJECT_STORAGE_URL

		if ph.URL != "" {
			ph.URL = "https://" + base + "/" + ph.URL
		}
		photos = append(photos, ph)

	}

	return photos
}

func (e *EntityPhotoRepositoryImplementation) FindAllEntityPhotosForEntity(entity string, entityID uint) []response.EntityPhotoResponse {
	var entityPhotos []models.EntityPhoto
	result := e.Db.Model(&models.EntityPhoto{}).Where("entity_id = ? AND entity_type = ?", entityID, entity).Find(&entityPhotos)
	if result.Error != nil {
		return []response.EntityPhotoResponse{}
	}

	var photos []response.EntityPhotoResponse
	var ph response.EntityPhotoResponse
	var photoModel models.Photo
	for _, entityPhoto := range entityPhotos {

		photo := e.Db.Model(&models.Photo{}).Where("id = ?", entityPhoto.PhotoID).First(&photoModel)

		if photo.Error == nil {
			ph.ID = entityPhoto.PhotoID
			ph.Photo.URL = entityPhoto.Photo.URL
		}

		ph.ID = photoModel.ID
		ph.Photo.URL = photoModel.URL

		// load config
		env, err := config.LoadConfig(".")
		if err != nil {
			fmt.Printf("error: %v", err)
		}

		base := env.OBJECT_STORAGE_URL

		if ph.Photo.URL != "" {
			ph.Photo.URL = "https://" + base + "/" + ph.Photo.URL
		}
		photos = append(photos, ph)

	}

	return photos
}
func (e *EntityPhotoRepositoryImplementation) AddPhotoToEntity(photoID uint, entity string, entityID uint) response.PhotoResponse {

	entityPhoto := models.EntityPhoto{
		PhotoID:    photoID,
		EntityID:   entityID,
		EntityType: entity,
	}

	result := e.Db.Create(&entityPhoto)
	if result.Error != nil {
		return response.PhotoResponse{}
	}

	return response.PhotoResponse{
		ID:  entityPhoto.PhotoID,
		URL: entityPhoto.Photo.URL,
	}
}
