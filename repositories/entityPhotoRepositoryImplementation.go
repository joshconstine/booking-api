package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

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
	result := e.Db.Where("entity_id = ? AND entity_type = ?", entityID, entity).Find(&entityPhotos)
	if result.Error != nil {
		return []response.PhotoResponse{}
	}

	var photos []response.PhotoResponse
	var ph response.PhotoResponse
	for _, entityPhoto := range entityPhotos {
		photo := e.Db.Where("id = ?", entityPhoto.PhotoID).First(&models.Photo{})
		if photo.Error == nil {
			ph.ID = entityPhoto.PhotoID
			ph.URL = entityPhoto.Photo.URL
		}
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
