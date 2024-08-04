package request

import (
	"booking-api/data/response"
	"booking-api/models"
	"gorm.io/gorm"
	"time"
)

type CreateRentalStep1Params struct {
	RentalID            uint
	Name                string
	Address             string
	Description         string
	Bedrooms            uint
	Bathrooms           float64
	Guests              uint
	AllowInstantBooking bool
	AllowPets           bool
	ParentProperty      bool
	Amenities           []response.AmenityResponse
	AccountID           uint
	Success             bool
}

type CreateRentalStep1Errors struct {
	Name        string
	Address     string
	Bedrooms    string
	Bathrooms   string
	Guests      string
	Description string
}

func (rental *CreateRentalStep1Params) MapCreateRentalStep1ToRental() models.Rental {

	fivePm := time.Date(2024, 0, 0, 17, 0, 0, 0, time.UTC)

	elevenAm := time.Date(2025, 0, 0, 11, 0, 0, 0, time.UTC)

	rentalModel := models.Rental{
		Name:             rental.Name,
		LocationID:       1,
		Bedrooms:         rental.Bedrooms,
		Bathrooms:        rental.Bathrooms,
		AccountID:        rental.AccountID,
		Description:      rental.Description,
		RentalStatus:     models.RentalStatus{},
		BookingCostItems: []models.EntityBookingCost{},
		BookingDurationRule: models.EntityBookingDurationRule{
			MinimumDuration: 1,
			MaximumDuration: 30,
			BookingBuffer:   1,
			StartTime:       fivePm,
			EndTime:         elevenAm,
		},
		BookingRule: models.EntityBookingRule{
			AdvertiseAtAllLocations: true,
			AllowPets:               true,
			AllowInstantBooking:     true,
			OfferEarlyCheckIn:       false,
		},
		Amenities:                  []models.Amenity{},
		RentalRooms:                []models.RentalRoom{},
		BookingCostItemAdjustments: []models.EntityBookingCostAdjustment{},
		BookingDocuments:           []models.EntityBookingDocument{},
		BookingRequests:            []models.EntityBookingPermission{},
		Timeblocks:                 []models.EntityTimeblock{},
		Bookings:                   []models.EntityBooking{},
		EntityPhotos: []models.EntityPhoto{
			{
				Photo: models.Photo{},
			},
		},
	}

	for _, amenity := range rental.Amenities {
		rentalModel.Amenities = append(rentalModel.Amenities, models.Amenity{
			Model: gorm.Model{
				ID: amenity.ID,
			},
		})
	}

	return rentalModel
}
