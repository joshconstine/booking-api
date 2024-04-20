package request

import (
	"booking-api/models"
	"time"
)

type CreateRentalRequest struct {
	Name        string  `json: validate:"required"`
	LocationID  uint    `json: validate:"required"`
	AccountID   uint    `json: validate:"required"`
	Bedrooms    uint    `json: validate:"required"`
	Bathrooms   uint    `json: validate:"required"`
	NightlyRate float64 `json: validate:"required"`
	Description string  `json: validate:"required"`
}

func (rental *CreateRentalRequest) MapCreateRentalRequestToRental() models.Rental {

	fivePm := time.Date(2024, 0, 0, 17, 0, 0, 0, time.UTC)

	elevenAm := time.Date(2025, 0, 0, 11, 0, 0, 0, time.UTC)

	rentalModel := models.Rental{
		Name:         rental.Name,
		LocationID:   rental.LocationID,
		Bedrooms:     rental.Bedrooms,
		Bathrooms:    rental.Bathrooms,
		AccountID:    rental.AccountID,
		Description:  rental.Description,
		RentalStatus: models.RentalStatus{},
		BookingCostItems: []models.EntityBookingCost{
			{
				Amount:            rental.NightlyRate,
				TaxRateID:         1,
				BookingCostTypeID: 1,
			},
		},
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
			AllowInstantBooking:     false,
			OfferEarlyCheckIn:       false,
		},
		Amenities:                  []models.Amenity{},
		RentalRooms:                []models.RentalRoom{},
		BookingCostItemAdjustments: []models.EntityBookingCostAdjustment{},
		BookingDocuments:           []models.EntityBookingDocument{},
		BookingRequests:            []models.EntityBookingRequest{},
		Timeblocks:                 []models.EntityTimeblock{},
		Bookings:                   []models.EntityBooking{},
		EntityPhotos:               []models.EntityPhoto{},
	}
	return rentalModel
}
