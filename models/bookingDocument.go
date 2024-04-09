package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BookingDocument struct {
	gorm.Model
	BookingID         string
	EntityBookingID   uint
	RequiresSignature bool
	Signed            bool
	Note              string
	DocumentID        uint
	Document          Document
	Booking           Booking
}

func (b *BookingDocument) TableName() string {
	return "booking_documents"
}

func (b *BookingDocument) MapBookingDocumentToResponse() response.BookingDocumentResponse {

	documentResponse := response.BookingDocumentResponse{
		ID:                b.ID,
		RequiresSignature: b.RequiresSignature,
		Note:              b.Note,
	}

	documentResponse.Document = b.Document.MapDocumentToResponse()

	return documentResponse

}
