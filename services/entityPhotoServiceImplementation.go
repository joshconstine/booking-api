package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type EntityPhotoServiceImplementation struct {
	EntityPhotoRepository repositories.EntityPhotoRepository
	Validate              *validator.Validate
}

func NewEntityPhotoServiceImplementation(entityPhotoRepository repositories.EntityPhotoRepository, validate *validator.Validate) EntityPhotoService {
	return &EntityPhotoServiceImplementation{
		EntityPhotoRepository: entityPhotoRepository,
		Validate:              validate,
	}
}

func (t EntityPhotoServiceImplementation) FindAllForEntity(entity string, entityID uint) []response.PhotoResponse {
	result := t.EntityPhotoRepository.FindAllForEntity(entity, entityID)

	return result
}

func (t EntityPhotoServiceImplementation) FindAllEntityPhotosForEntity(entity string, entityID uint) []response.EntityPhotoResponse {
	result := t.EntityPhotoRepository.FindAllEntityPhotosForEntity(entity, entityID)

	return result
}

func (t EntityPhotoServiceImplementation) AddPhotoToEntity(photoID uint, entity string, entityID uint) response.PhotoResponse {
	result := t.EntityPhotoRepository.AddPhotoToEntity(photoID, entity, entityID)

	return result
}
