package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

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
	result := r.Db.Model(&models.Rental{}).Preload("Location").Preload("Amenities").Preload("Photos").Preload("RentalRooms").Preload("BookingCostItems").Preload("BookingDurationRule").Preload("Timeblocks").Find(&rental)
	if result.Error != nil {
		return response.RentalInformationResponse{}
	}

	return rental.MapRentalToInformationResponse()

}

func (r *RentalRepositoryImplementation) Create(rental models.Rental) models.Rental {
	result := r.Db.Create(&rental)
	if result.Error != nil {
		return models.Rental{}
	}

	return rental
}
