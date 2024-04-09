package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type EntityBookingDocumentRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityBookingDocumentRepositoryImplementation(db *gorm.DB) *EntityBookingDocumentRepositoryImplementation {
	return &EntityBookingDocumentRepositoryImplementation{Db: db}
}

func (e *EntityBookingDocumentRepositoryImplementation) FindEntityBookingDocumentsByEntityIDAndEntityType(entityID uint, entityType string) []response.EntityBookingDocumentResponse {
	var entityBookingDocuments []models.EntityBookingDocument
	e.Db.Where("entity_id = ? AND entity_type = ?", entityID, entityType).Preload("Document").Find(&entityBookingDocuments)

	var result []response.EntityBookingDocumentResponse

	for _, entityBookingDocument := range entityBookingDocuments {
		result = append(result, entityBookingDocument.MapEntityBookingDocumentToResponse())
	}

	return result
}

func (e *EntityBookingDocumentRepositoryImplementation) Create(request request.CreateEntityBookingDocumentRequest) (response.EntityBookingDocumentResponse, error) {
	entityBookingDocument := models.EntityBookingDocument{
		EntityID:          request.EntityID,
		EntityType:        request.EntityType,
		DocumentID:        request.DocumentID,
		RequiresSignature: request.RequiresSignature,
	}

	result := e.Db.Create(&entityBookingDocument)
	if result.Error != nil {
		return response.EntityBookingDocumentResponse{}, result.Error
	}

	return entityBookingDocument.MapEntityBookingDocumentToResponse(), nil
}
