package services

import (
	"booking-api/data/response"
	"mime/multipart"
)

type PhotoService interface {
	FindAll() []response.PhotoResponse
	AddPhoto(photo *multipart.File, header *multipart.FileHeader, entity string, entityID int) response.PhotoResponse
}
