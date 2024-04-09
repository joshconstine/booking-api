package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type EntityBookingDocumentRepository interface {
	FindEntityBookingDocumentsByEntityIDAndEntityType(entityID uint, entityType string) []response.EntityBookingDocumentResponse
	Create(request request.CreateEntityBookingDocumentRequest) (response.EntityBookingDocumentResponse, error)
}
