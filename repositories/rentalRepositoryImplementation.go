package repositories

import (
	resposes "booking-api/data/response"
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

func (r *RentalRepositoryImplementation) FindAll() []models.Rental {
	var rentals []models.Rental
	result := r.Db.Find(&rentals)
	if result.Error != nil {
		return []models.Rental{}
	}

	return rentals
}

func (r *RentalRepositoryImplementation) FindById(id uint) models.Rental {
	var rental models.Rental
	result := r.Db.Where("id = ?", id).First(&rental)
	if result.Error != nil {
		return models.Rental{}
	}

	return rental
}

func (r *RentalRepositoryImplementation) Create(rental models.Rental) models.Rental {
	result := r.Db.Create(&rental)
	if result.Error != nil {
		return models.Rental{}
	}

	return rental
}

func (r *RentalRepositoryImplementation) GetInformationForRental(id uint) (rentalInformation resposes.RentalInformationResponse) {
	var timeblocks []models.Timeblock

	rental := r.FindById(id)

	rentalInformation.RentalID = rental.ID
	rentalInformation.Name = rental.Name
	rentalInformation.LocationID = rental.LocationID
	rentalInformation.LocationName = rental.Location.Name
	rentalInformation.RentalIsClean = false
	rentalInformation.Thumbnail = "https://via.placeholder.com/150"

	// timeblocks := r.TimeblockRepository.FindByEntity("rentals", id)

	rentalInformation.Timeblocks = timeblocks

	return
}
