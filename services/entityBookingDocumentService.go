package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingDocumentService interface {
	FindEntityBookingDocumentsForEntity(entityID uint, entityType string) []response.EntityBookingDocumentResponse
	CreateEntityBookingDocument(request request.CreateEntityBookingDocumentRequest) (response.EntityBookingDocumentResponse, error)
}
