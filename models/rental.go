package models

import (
	"booking-api/config"
	"booking-api/data/response"
	"fmt"

	"gorm.io/gorm"
)

type Rental struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	LocationID  uint    `gorm:"not null"`
	Bedrooms    uint    `gorm:"not null"`
	Bathrooms   float64 `gorm:"not null"`
	AccountID   uint    `gorm:"not null"`
	Description string

	Location     Location
	RentalStatus RentalStatus

	Amenities                  []Amenity                     `gorm:"many2many:rental_amenities;"`
	Timeblocks                 []EntityTimeblock             `gorm:"polymorphic:Entity"`
	EntityPhotos               []EntityPhoto                 `gorm:"polymorphic:Entity"`
	RentalRooms                []RentalRoom                  `gorm:"foreignKey:RentalID"`
	Bookings                   []EntityBooking               `gorm:"polymorphic:Entity"`
	BookingCostItems           []EntityBookingCost           `gorm:"polymorphic:Entity"`
	BookingCostItemAdjustments []EntityBookingCostAdjustment `gorm:"polymorphic:Entity"`
	BookingDocuments           []EntityBookingDocument       `gorm:"polymorphic:Entity"`
	BookingDurationRule        EntityBookingDurationRule     `gorm:"polymorphic:Entity"`
	BookingRule                EntityBookingRule             `gorm:"polymorphic:Entity"`
	Reviews                    []EntityReview                `gorm:"polymorphic:Entity"`

	BookingRequests []EntityBookingPermission `gorm:"polymorphic:Entity"`
}

func (r *Rental) TableName() string {
	return "rentals"
}

func (r *Rental) MapRentalsToResponse() response.RentalResponse {
	thumbnail := ""

	if len(r.EntityPhotos) > 0 {

		// load config
		env, err := config.LoadConfig(".")
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		base := env.OBJECT_STORAGE_URL
		thumbnail = r.EntityPhotos[0].Photo.URL

		if thumbnail != "" {
			thumbnail = "https://" + base + "/" + thumbnail
		}

	}

	return response.RentalResponse{
		ID:          r.ID,
		Name:        r.Name,
		Location:    r.Location.MapLocationToResponse(),
		Bedrooms:    r.Bedrooms,
		Bathrooms:   r.Bathrooms,
		Description: r.Description,
		Thumbnail:   thumbnail,
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

	for _, bookingCostItemAdjustment := range r.BookingCostItemAdjustments {
		response.BookingCostItemAdjustments = append(response.BookingCostItemAdjustments, bookingCostItemAdjustment.MapEntityBookingCostAdjustmentToResponse())
	}

	for _, timeblock := range r.Timeblocks {
		response.Timeblocks = append(response.Timeblocks, timeblock.MapTimeblockToResponse())
	}

	for _, document := range r.BookingDocuments {
		response.BookingDocuments = append(response.BookingDocuments, document.MapEntityBookingDocumentToResponse())
	}

	for _, inquiry := range r.BookingRequests {
		response.BookingRequests = append(response.BookingRequests, inquiry.MapEntityBookingPermissionToResponse())
	}

	response.BookingDurationRule = r.BookingDurationRule.MapEntityBookingDurationRuleToResponse()
	response.BookingRule = r.BookingRule.MapEntityBookingRuleToResponse()

	return response
}
