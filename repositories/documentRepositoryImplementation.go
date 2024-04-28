package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/pkg/objectStorage"
	"mime/multipart"

	"gorm.io/gorm"
)

type DocumentRepositoryImplementation struct {
	StorageClient *objectStorage.S3Client
	DB            *gorm.DB
}

func NewDocumentRepositoryImplementation(storageClient *objectStorage.S3Client, db *gorm.DB) DocumentRepository {
	return &DocumentRepositoryImplementation{StorageClient: storageClient, DB: db}
}

// FindAll retrieves all documents from the database
func (r *DocumentRepositoryImplementation) FindAll() []response.DocumentResponse {

	var documents []models.Document
	result := r.DB.Model(&documents).Find(&documents)

	if result.Error != nil {
		return []response.DocumentResponse{}
	}

	var documentResponses []response.DocumentResponse
	for _, document := range documents {
		documentResponses = append(documentResponses, document.MapDocumentToResponse())
	}

	return documentResponses

}

// FindById retrieves a document by its ID from the database
func (r *DocumentRepositoryImplementation) FindById(id uint) response.DocumentResponse {

	var document models.Document
	result := r.DB.Model(&models.Document{}).Where("id = ?", id).First(&document)

	if result.Error != nil {
		return response.DocumentResponse{}
	}

	return document.MapDocumentToResponse()

}

// AddDocument adds a document to the database
func (r *DocumentRepositoryImplementation) AddDocument(filePath string, document *multipart.File, fileExt string, name string) response.DocumentResponse {

	returnedUrl, err := r.StorageClient.UploadFile(document, filePath, fileExt)

	if err != nil {
		return response.DocumentResponse{}
	}

	var documentModel models.Document
	documentModel.URL = returnedUrl
	documentModel.Name = name

	result := r.DB.Create(&documentModel)

	if result.Error != nil {
		return response.DocumentResponse{}
	}

	return documentModel.MapDocumentToResponse()

}

// AddDocumentWithUrl adds a document to the database with a given URL
func (r *DocumentRepositoryImplementation) AddDocumentWithUrl(url string, name string) response.DocumentResponse {

	var documentModel models.Document
	documentModel.URL = url
	documentModel.Name = name

	result := r.DB.Create(&documentModel)

	if result.Error != nil {
		return response.DocumentResponse{}
	}

	return documentModel.MapDocumentToResponse()

}
