package request

import (
	"booking-api/models"
	"time"
)

type CreateBoatRequest struct {
	Name        string
	AccountID   uint
	NightlyRate float64
	Description string
	Thumbnail   string
}

func (boat *CreateBoatRequest) MapCreateBoatRequestToBoat(taxid uint) models.Boat {

	fivePm := time.Date(2024, 0, 0, 17, 0, 0, 0, time.UTC)

	elevenAm := time.Date(2025, 0, 0, 11, 0, 0, 0, time.UTC)

	boatModel := models.Boat{
		Name:      boat.Name,
		AccountID: boat.AccountID,
		Status:    models.BoatStatus{},
		BookingCostItems: []models.EntityBookingCost{
			{
				Amount:    boat.NightlyRate,
				TaxRateID: taxid,
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
		BookingCostItemAdjustments: []models.EntityBookingCostAdjustment{},
		BookingDocuments:           []models.EntityBookingDocument{},
		BookingRequests:            []models.EntityBookingPermission{},
		Timeblocks:                 []models.EntityTimeblock{},
		Bookings:                   []models.EntityBooking{},
		Photos: []models.EntityPhoto{
			{
				Photo: models.Photo{
					URL: boat.Thumbnail,
				},
			},
		},
	}
	return boatModel
}
