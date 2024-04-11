package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type EntityBookingDocumentServiceImplementation struct {
	EntityBookingDocumentRepository repositories.EntityBookingDocumentRepository
}

func NewEntityBookingDocumentServiceImplementation(entityBookingDocumentRepository repositories.EntityBookingDocumentRepository) EntityBookingDocumentService {
	return &EntityBookingDocumentServiceImplementation{EntityBookingDocumentRepository: entityBookingDocumentRepository}
}

func (e *EntityBookingDocumentServiceImplementation) FindEntityBookingDocumentsForEntity(entityID uint, entityType string) []response.EntityBookingDocumentResponse {
	return e.EntityBookingDocumentRepository.FindEntityBookingDocumentsByEntityIDAndEntityType(entityID, entityType)
}

func (e *EntityBookingDocumentServiceImplementation) CreateEntityBookingDocument(request request.CreateEntityBookingDocumentRequest) (response.EntityBookingDocumentResponse, error) {
	return e.EntityBookingDocumentRepository.Create(request)
}
