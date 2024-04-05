package services

import (
	"booking-api/constants"
	"booking-api/data/response"
	"booking-api/repositories"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PhotoServiceImplementation struct {
	PhotoRepository repositories.PhotoRepository
	Validate        *validator.Validate
}

func BuildFilePath(entity string, entityID int, fileExt string) string {

	newFilename := uuid.New().String() + fileExt
	var location string

	if entity == constants.REANTAL_ENTITY {
		location = filepath.Join(constants.RENTAL_PHOTO_EXTENSION, strconv.Itoa(entityID))

	} else if entity == constants.BOAT_ENTITY {
		location = filepath.Join(constants.BOAT_PHOTO_EXTENSION, strconv.Itoa(entityID))
	}

	return filepath.Join(location, newFilename)

}

func NewPhotoServiceImplementation(photoRepository repositories.PhotoRepository, validate *validator.Validate) PhotoService {
	return &PhotoServiceImplementation{PhotoRepository: photoRepository, Validate: validate}
}

func (service *PhotoServiceImplementation) FindAll() []response.PhotoResponse {
	return service.PhotoRepository.FindAll()
}

func (service *PhotoServiceImplementation) AddPhoto(file *multipart.File, header *multipart.FileHeader, entity string, entityID int) response.PhotoResponse {

	fileExt := filepath.Ext(header.Filename)
	filePath := BuildFilePath(entity, entityID, fileExt)
	return service.PhotoRepository.AddPhoto(filePath, file, fileExt)
}
