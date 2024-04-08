package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type BookingDocument struct {
	gorm.Model
	BookingID         uint
	RequiresSignature bool
	Signed            bool
	Note              string
	Document          Document
	Booking           Booking
}

func (b *BookingDocument) TableName() string {
	return "booking_documents"
}

func (b *BookingDocument) MapBookingDocumentToResponse() response.BookingDocumentResponse {

	documentResponse := response.BookingDocumentResponse{
		ID:        b.ID,
		BookingID: b.BookingID,

		RequiresSignature: b.RequiresSignature,
		Note:              b.Note,
	}

	return documentResponse

}
