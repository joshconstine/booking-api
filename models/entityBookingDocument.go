package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityBookingDocument struct {
	gorm.Model
	EntityID          uint   `gorm:"index:idx_entity_document_type,unique"`
	EntityType        string `gorm:"index:idx_entity_document_type,unique"`
	DocumentID        uint   `gorm:"index:idx_entity_document_type,unique"`
	RequiresSignature bool
	Document          Document
}

func (e *EntityBookingDocument) TableName() string {
	return "entity_booking_documents"
}

func (e *EntityBookingDocument) MapEntityBookingDocumentToResponse() response.EntityBookingDocumentResponse {

	result := response.EntityBookingDocumentResponse{
		ID:                e.ID,
		RequiresSignature: e.RequiresSignature,
	}

	result.Document = e.Document.MapDocumentToResponse()

	return result

}
