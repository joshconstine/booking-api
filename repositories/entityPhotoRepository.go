package repositories

import (
	"booking-api/data/response"
)

type EntityPhotoRepository interface {
	FindAllForEntity(entity string, entityID uint) []response.PhotoResponse
	AddPhotoToEntity(photoID uint, entity string, entityID uint) response.PhotoResponse
	// RemovePhotoFromEntity(photoID uint, entity string, entityID uint) response.PhotoResponse
}
