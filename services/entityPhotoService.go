package services

import (
	"booking-api/data/response"
)

type EntityPhotoService interface {
	FindAllForEntity(entity string, entityID uint) []response.PhotoResponse
	AddPhotoToEntity(photoID uint, entity string, entityID uint) response.PhotoResponse
	FindAllEntityPhotosForEntity(entity string, entityID uint) []response.EntityPhotoResponse
}
