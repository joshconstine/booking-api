package request

import (
	"booking-api/constants"
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
	Thumbnail   string
}

func (rental *CreateRentalRequest) MapCreateRentalRequestToRental(taxid uint) models.Rental {

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
				BookingCostTypeID: constants.BOOKING_COST_TYPE_RENTAL_COST_ID,
				TaxRateID:         taxid,
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
				Photo: models.Photo{
					URL: rental.Thumbnail,
				},
			},
		},
	}
	return rentalModel
}
