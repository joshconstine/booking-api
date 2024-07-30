package repositories

import (
	"booking-api/data/response"
	"mime/multipart"
)

type PhotoRepository interface {
	FindAll() []response.PhotoResponse
	AddPhoto(filePath string, photo *multipart.File, fileExt string) response.PhotoResponse
}
