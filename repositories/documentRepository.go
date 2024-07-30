package repositories

import (
	"booking-api/data/response"
	"mime/multipart"
)

type DocumentRepository interface {
	FindAll() []response.DocumentResponse
	FindById(id uint) response.DocumentResponse
	AddDocument(filePath string, document *multipart.File, fileExt string, name string) response.DocumentResponse
	AddDocumentWithUrl(url string, name string) response.DocumentResponse
}
