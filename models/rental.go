package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Rental struct {
	gorm.Model
	Name         string
	LocationID   uint
	Bedrooms     int
	Bathrooms    int
	Description  string
	Location     Location
	RentalStatus RentalStatus
	AccountID    uint `gorm:"not null"`

	Amenities           []Amenity                 `gorm:"many2many:rental_amenities;"`
	Timeblocks          []EntityTimeblock         `gorm:"polymorphic:Entity"`
	EntityPhotos        []EntityPhoto             `gorm:"polymorphic:Entity"`
	RentalRooms         []RentalRoom              `gorm:"foreignKey:RentalID"`
	Bookings            []EntityBooking           `gorm:"polymorphic:Entity"`
	BookingCostItems    []EntityBookingCost       `gorm:"polymorphic:Entity"`
	BookingDocuments    []EntityBookingDocument   `gorm:"polymorphic:Entity"`
	BookingDurationRule EntityBookingDurationRule `gorm:"polymorphic:Entity"`
	BookingRule         EntityBookingRule         `gorm:"polymorphic:Entity"`

	BookingRequests []EntityBookingRequest `gorm:"polymorphic:Entity"`
}

func (r *Rental) TableName() string {
	return "rentals"
}

func (r *Rental) MapRentalsToResponse() response.RentalResponse {
	return response.RentalResponse{
		ID:          r.ID,
		Name:        r.Name,
		Location:    r.Location.MapLocationToResponse(),
		Bedrooms:    r.Bedrooms,
		Bathrooms:   r.Bathrooms,
		Description: r.Description,
	}
}

func (r *Rental) MapRentalToInformationResponse() response.RentalInformationResponse {
	var response response.RentalInformationResponse
	// Enable detailed log of operations

	response.ID = r.ID
	response.Name = r.Name
	response.Bedrooms = r.Bedrooms
	response.Bathrooms = r.Bathrooms
	response.Description = r.Description
	response.Location = r.Location.MapLocationToResponse()
	response.RentalStatus = r.RentalStatus.MapRentalStatusToResponse()

	for _, amenity := range r.Amenities {
		response.Amenities = append(response.Amenities, amenity.MapAmenityToResponse())
	}

	for _, photo := range r.EntityPhotos {
		response.Photos = append(response.Photos, photo.MapEntityPhotoToResponse())
	}

	for _, rentalRoom := range r.RentalRooms {
		response.RentalRooms = append(response.RentalRooms, rentalRoom.MapRentalRoomToResponse())
	}

	for _, booking := range r.Bookings {
		response.Bookings = append(response.Bookings, booking.MapEntityBookingToResponse())
	}

	for _, bookingCostItem := range r.BookingCostItems {
		response.BookingCostItems = append(response.BookingCostItems, bookingCostItem.MapEntityBookingCostToResponse())
	}
	for _, timeblock := range r.Timeblocks {
		response.Timeblocks = append(response.Timeblocks, timeblock.MapTimeblockToResponse())
	}

	for _, document := range r.BookingDocuments {
		response.BookingDocuments = append(response.BookingDocuments, document.MapEntityBookingDocumentToResponse())
	}

	for _, inquiry := range r.BookingRequests {
		response.BookingRequests = append(response.BookingRequests, inquiry.MapEntityBookingRequestToResponse())
	}

	response.BookingDurationRule = r.BookingDurationRule.MapEntityBookingDurationRuleToResponse()
	response.BookingRule = r.BookingRule.MapEntityBookingRuleToResponse()

	return response
}
