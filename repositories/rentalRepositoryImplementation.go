package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"time"

	"gorm.io/gorm"
)

type RentalRepositoryImplementation struct {
	Db                  *gorm.DB
	TimeblockRepository TimeblockRepository
}

func NewRentalRepositoryImplementation(db *gorm.DB, timeblockRepository TimeblockRepository) RentalRepository {
	return &RentalRepositoryImplementation{Db: db, TimeblockRepository: timeblockRepository}
}

func (r *RentalRepositoryImplementation) FindAll() []response.RentalResponse {
	var rentals []models.Rental
	result := r.Db.Model(&models.Rental{}).Preload("Location").Find(&rentals)
	if result.Error != nil {
		return []response.RentalResponse{}
	}

	var rentalResponses []response.RentalResponse
	for _, rental := range rentals {
		rentalResponses = append(rentalResponses, rental.MapRentalsToResponse())
	}

	return rentalResponses
}

func (r *RentalRepositoryImplementation) FindById(id uint) response.RentalInformationResponse {
	var rental models.Rental

	// Adjusting the Joins method to correctly reference the intermediary table and both sides of the many-to-many relationship
	result := r.Db.Model(&models.Rental{}).
		Joins("LEFT JOIN rental_amenities on rental_amenities.rental_id = rentals.id").
		Joins("LEFT JOIN amenities on amenities.id = rental_amenities.amenity_id").
		Where("rentals.id = ?", id).
		Preload("Location").
		Preload("RentalStatus").
		Preload("Amenities"). // This might still be necessary to preload the related Amenities correctly
		Preload("Amenities.AmenityType").
		Preload("EntityPhotos").
		Preload("EntityPhotos.Photo").
		Preload("RentalRooms.Beds").
		Preload("RentalRooms.RoomType").
		Preload("RentalRooms.Photos").
		Preload("RentalRooms.Photos.Photo").
		Preload("BookingCostItems.BookingCostType").
		Preload("BookingCostItems.TaxRate").
		Preload("BookingDurationRule").
		Preload("BookingRule").
		Preload("Timeblocks").
		Preload("Bookings").
		Preload("BookingDocuments.Document").
		Preload("BookingRequests.InquiryStatus").
		Preload("BookingCostItemAdjustments.BookingCostType").
		Preload("BookingCostItemAdjustments.TaxRate").
		First(&rental)

	if result.Error != nil {
		return response.RentalInformationResponse{}
	}

	return rental.MapRentalToInformationResponse()
}

func (r *RentalRepositoryImplementation) Create(rental request.CreateRentalRequest) (response.RentalResponse, error) {
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

	result := r.Db.Create(&rentalModel)
	if result.Error != nil {
		return response.RentalResponse{}, result.Error
	}

	return rentalModel.MapRentalsToResponse(), nil
}
