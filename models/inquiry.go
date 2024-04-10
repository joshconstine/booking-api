package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Inquiry struct {
	gorm.Model
	UserID          uint `gorm:"not null"`
	BookingID       string
	Note            string
	NumGuests       int
	User            User
	BookingRequests []EntityBookingRequest
}

func (inquiry *Inquiry) TableName() string {
	return "inquiries"
}

func (inquiry *Inquiry) MapInquiryToResponse() response.InquiryResponse {
	inquiryResponse := response.InquiryResponse{
		ID:        inquiry.ID,
		Note:      inquiry.Note,
		NumGuests: inquiry.NumGuests,
		User:      inquiry.User.MapUserToResponse(),
	}

	for _, entityInquiry := range inquiry.BookingRequests {
		inquiryResponse.BookingRequests = append(inquiryResponse.BookingRequests, entityInquiry.MapEntityBookingRequestToResponse())
	}

	return inquiryResponse

}
