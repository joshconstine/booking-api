package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/pkg/database"
	"log"

	"gorm.io/gorm"
)

type RentalRepositoryImplementation struct {
	Db                  *gorm.DB
	TimeblockRepository EntityTimeblockRepository
}

func NewRentalRepositoryImplementation(db *gorm.DB, timeblockRepository EntityTimeblockRepository) RentalRepository {
	return &RentalRepositoryImplementation{Db: db, TimeblockRepository: timeblockRepository}
}

func (r *RentalRepositoryImplementation) FindAll() []response.RentalResponse {
	var rentals []models.Rental
	result := r.Db.Model(&models.Rental{}).Preload("Location").Preload("EntityPhotos.Photo").Find(&rentals)
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

	//TODO Add the tax rate to the create request to remove the need for this query
	taxRates := []models.TaxRate{}

	result := database.Instance.Model(&models.TaxRate{}).Find(&taxRates)
	if result.Error != nil {
		log.Println(result.Error)
	}

	var shortTermTaxRate models.TaxRate
	for _, taxRate := range taxRates {
		if taxRate.Name == "Short Term Rental Tax" {
			shortTermTaxRate = taxRate
		}

	}

	rentalModel := rental.MapCreateRentalRequestToRental(shortTermTaxRate.ID)
	result = r.Db.Create(&rentalModel)
	if result.Error != nil {
		return response.RentalResponse{}, result.Error
	}

	return rentalModel.MapRentalsToResponse(), nil
}

func (r *RentalRepositoryImplementation) CreateStep1(rental request.CreateRentalStep1Params) (response.RentalResponse, error) {

	rentalModel := rental.MapCreateRentalStep1ToRental()
	result := r.Db.Create(&rentalModel)
	if result.Error != nil {
		return response.RentalResponse{}, result.Error
	}

	return rentalModel.MapRentalsToResponse(), nil
}

func (r *RentalRepositoryImplementation) Update(rental request.UpdateRentalRequest) (response.RentalResponse, error) {
	var rentalToUpdate models.Rental
	result := r.Db.Where("id = ?", rental.ID).First(&rentalToUpdate)
	if result.Error != nil {
		return response.RentalResponse{}, result.Error
	}

	rentalToUpdate.Name = rental.Name
	rentalToUpdate.Bedrooms = rental.Bedrooms
	rentalToUpdate.Bathrooms = rental.Bathrooms
	rentalToUpdate.Description = rental.Description

	result = r.Db.Save(&rentalToUpdate)
	if result.Error != nil {
		return response.RentalResponse{}, result.Error
	}

	return rentalToUpdate.MapRentalsToResponse(), nil
}

func (r *RentalRepositoryImplementation) UpdateRental(rental request.CreateRentalStep1Params) (response.RentalResponse, error) {
	var rentalToUpdate models.Rental
	result := r.Db.Where("id = ?", rental.RentalID).Preload("BookingRule").First(&rentalToUpdate)
	if result.Error != nil {
		return response.RentalResponse{}, result.Error
	}

	rentalToUpdate.Name = rental.Name
	rentalToUpdate.Bedrooms = uint(rental.Bedrooms)
	rentalToUpdate.Bathrooms = rental.Bathrooms
	rentalToUpdate.Description = rental.Description
	rentalToUpdate.Address = rental.Address
	rentalToUpdate.Guests = rental.Guests

	if rentalToUpdate.BookingRule.AllowInstantBooking != rental.AllowInstantBooking || rentalToUpdate.BookingRule.AllowPets != rental.AllowPets {
		rentalToUpdate.BookingRule.AllowInstantBooking = rental.AllowInstantBooking
		rentalToUpdate.BookingRule.AllowPets = rental.AllowPets
		updateResult := r.Db.Save(&rentalToUpdate.BookingRule)
		if updateResult.Error != nil {
			return response.RentalResponse{}, updateResult.Error
		}
	}

	rentalToUpdate.Amenities = []models.Amenity{}
	for _, amenity := range rental.Amenities {
		rentalToUpdate.Amenities = append(rentalToUpdate.Amenities, models.Amenity{
			Model: gorm.Model{
				ID: amenity.ID,
			},
		})

	}

	result = r.Db.Save(&rentalToUpdate)
	if result.Error != nil {
		return response.RentalResponse{}, result.Error
	}

	return rentalToUpdate.MapRentalsToResponse(), nil
}

func (r *RentalRepositoryImplementation) FindAllIDs() []uint {
	var rentalIDs []uint
	result := r.Db.Model(&models.Rental{}).Pluck("id", &rentalIDs)
	if result.Error != nil {
		return []uint{}
	}

	return rentalIDs
}
