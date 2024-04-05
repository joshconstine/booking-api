package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type PhotoServiceImplementation struct {
	PhotoRepository repositories.PhotoRepository
	Validate        *validator.Validate
}

func NewPhotoServiceImplementation(photoRepository repositories.PhotoRepository, validate *validator.Validate) PhotoService {
	return &PhotoServiceImplementation{PhotoRepository: photoRepository, Validate: validate}
}

func (service *PhotoServiceImplementation) FindAll() []response.PhotoResponse {
	return service.PhotoRepository.FindAll()
}
