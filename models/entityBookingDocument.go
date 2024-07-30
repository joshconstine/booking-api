package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type EntityBookingDocument struct {
	gorm.Model
	EntityID          uint   `gorm:"primaryKey"`
	EntityType        string `gorm:"primaryKey"`
	DocumentID        uint   `gorm:"not null"`
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

func (e *EntityBookingDocument) MapEntityBookingDocumentToBookingDocument() BookingDocument {

	result := BookingDocument{
		RequiresSignature: e.RequiresSignature,
		Note:              "",
		DocumentID:        e.DocumentID,
	}

	return result
}
