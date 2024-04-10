package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Boat struct {
	gorm.Model
	Name      string
	Occupancy uint       `gorm:"not null"`
	MaxWeight uint       `gorm:"not null"`
	AccountID uint       `gorm:"not null"`
	Status    BoatStatus `gorm:"not null"`

	Timeblocks                 []EntityTimeblock             `gorm:"polymorphic:Entity"`
	Photos                     []EntityPhoto                 `gorm:"polymorphic:Entity"`
	Bookings                   []EntityBooking               `gorm:"polymorphic:Entity"`
	BookingCostItems           []EntityBookingCost           `gorm:"polymorphic:Entity"`
	BookingCostItemAdjustments []EntityBookingCostAdjustment `gorm:"polymorphic:Entity"`
	BookingDocuments           []EntityBookingDocument       `gorm:"polymorphic:Entity"`
	BookingDurationRule        EntityBookingDurationRule     `gorm:"polymorphic:Entity; not null"`
	BookingRule                EntityBookingRule             `gorm:"polymorphic:Entity; not null"`

	BookingRequests []EntityBookingRequest `gorm:"polymorphic:Entity"`
}

func (b *Boat) TableName() string {
	return "boats"
}

func (b *Boat) MapBoatToResponse() response.BoatResponse {
	result := response.BoatResponse{
		ID:        b.ID,
		Name:      b.Name,
		Occupancy: b.Occupancy,
		MaxWeight: b.MaxWeight,
	}

	for _, timeBlock := range b.Timeblocks {
		result.Timeblocks = append(result.Timeblocks, timeBlock.MapTimeblockToResponse())
	}

	for _, photo := range b.Photos {
		result.Photos = append(result.Photos, photo.MapEntityPhotoToResponse())
	}

	result.BookingDurationRule = b.BookingDurationRule.MapEntityBookingDurationRuleToResponse()
	result.BookingRule = b.BookingRule.MapEntityBookingRuleToResponse()
	return result
}

func (b *Boat) MapBoatToInformationResponse() response.BoatInformationResponse {
	result := response.BoatInformationResponse{
		ID:        b.ID,
		Name:      b.Name,
		Occupancy: b.Occupancy,
		MaxWeight: b.MaxWeight,
	}

	for _, timeBlock := range b.Timeblocks {
		result.Timeblocks = append(result.Timeblocks, timeBlock.MapTimeblockToResponse())
	}

	for _, photo := range b.Photos {
		result.Photos = append(result.Photos, photo.MapEntityPhotoToResponse())
	}

	for _, booking := range b.Bookings {
		result.Bookings = append(result.Bookings, booking.MapEntityBookingToResponse())
	}

	for _, bookingCostItem := range b.BookingCostItems {
		result.BookingCostItems = append(result.BookingCostItems, bookingCostItem.MapEntityBookingCostToResponse())
	}

	for _, bookingCostItemAdjustment := range b.BookingCostItemAdjustments {
		result.BookingCostItemAdjustments = append(result.BookingCostItemAdjustments, bookingCostItemAdjustment.MapEntityBookingCostAdjustmentToResponse())
	}

	for _, bookingDocument := range b.BookingDocuments {
		result.BookingDocuments = append(result.BookingDocuments, bookingDocument.MapEntityBookingDocumentToResponse())
	}

	for _, bookingRequest := range b.BookingRequests {
		result.BookingRequests = append(result.BookingRequests, bookingRequest.MapEntityBookingRequestToResponse())
	}

	result.Status = b.Status.MapBoatStatusToResponse()
	result.BookingDurationRule = b.BookingDurationRule.MapEntityBookingDurationRuleToResponse()
	result.BookingRule = b.BookingRule.MapEntityBookingRuleToResponse()
	return result
}
