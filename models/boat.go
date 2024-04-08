package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Boat struct {
	gorm.Model
	Name                string
	Occupancy           int
	MaxWeight           int
	Timeblocks          []Timeblock               `gorm:"polymorphic:Entity"`
	Photos              []EntityPhoto             `gorm:"polymorphic:Entity"`
	Bookings            []EntityBooking           `gorm:"polymorphic:Entity"`
	BookingCostItems    []EntityBookingCost       `gorm:"polymorphic:Entity"`
	BookingDurationRule EntityBookingDurationRule `gorm:"polymorphic:Entity"`
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

	for _, booking := range b.Bookings {
		result.Bookings = append(result.Bookings, booking.MapEntityBookingToResponse())
	}

	for _, bookingCostItem := range b.BookingCostItems {
		result.BookingCostItems = append(result.BookingCostItems, bookingCostItem.MapEntityBookingCostToResponse())
	}

	result.BookingDurationRule = b.BookingDurationRule.MapEntityBookingDurationRuleToResponse()
	return result
}
