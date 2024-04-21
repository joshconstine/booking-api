package models

import (
	"booking-api/data/response"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID              string `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          string         ` gorm:"not null"`
	BookingStatusID uint           `gorm:"not null; default:1"`
	InquiryID       uint
	Inquiry         Inquiry
	// User            User `gorm:"not nill; foreignKey:UserID"`
	BookingStatus BookingStatus
	Details       BookingDetails `gorm:"not null"`
	CostItems     []BookingCostItem
	Payments      []BookingPayment
	Documents     []BookingDocument
	Entities      []EntityBooking
}

func (b *Booking) TableName() string {
	return "bookings"
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}
func (b *Booking) BeforeCreate(db *gorm.DB) error {
	rand := RandStringBytesMask(10)

	b.ID = rand

	return nil

}

func (b *Booking) MapBookingToResponse() response.BookingResponse {

	response := response.BookingResponse{
		ID: b.ID,
		// UserID:    b.UserID,
		// StatusID:  b.StatusID,
		// CreatedAt: b.CreatedAt,
		// UpdatedAt: b.UpdatedAt,
	}

	// for _, costItem := range b.CostItems {
	// 	response.CostItems = append(response.CostItems, costItem.MapBookingCostItemToResponse())
	// }

	// for _, payment := range b.Payments {
	// 	response.Payments = append(response.Payments, payment.MapBookingPaymentToResponse())
	// }

	// for _, document := range b.Documents {
	// 	response.Documents = append(response.Documents, document.MapBookingDocumentToResponse())
	// }

	return response
}

func (b *Booking) MapBookingToInformationResponse() response.BookingInformationResponse {

	response := response.BookingInformationResponse{
		ID: b.ID,
		Customer: response.UserResponse{
			UserID: b.UserID,
		},
	}

	response.Status = b.BookingStatus.MapBookingStatusToResponse()
	response.Details = b.Details.MapBookingDetailsToResponse()

	for _, costItem := range b.CostItems {
		response.CostItems = append(response.CostItems, costItem.MapBookingCostItemToResponse())
	}

	for _, payment := range b.Payments {
		response.Payments = append(response.Payments, payment.MapBookingPaymentToResponse())
	}

	for _, document := range b.Documents {
		response.Documents = append(response.Documents, document.MapBookingDocumentToResponse())
	}

	for _, entity := range b.Entities {
		response.Entities = append(response.Entities, entity.MapEntityBookingToResponse())
	}

	// response.Customer = b.User.MapUserToResponse()

	//log the booking details
	fmt.Println(response)
	return response
}
