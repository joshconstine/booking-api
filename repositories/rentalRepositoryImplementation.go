package repositories

import (
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
	result := r.Db.Model(&models.Rental{}).Preload("Location").Preload("RentalStatus").Preload("RentalRooms").Preload("Photos").Preload("BookingDurationRule").Preload("Bookings").Preload("BookingCostItems").Find(&rentals)
	if result.Error != nil {
		return []models.Rental{}
	}

	return rentals
}

func (r *RentalRepositoryImplementation) FindById(id uint) models.Rental {
	var rental models.Rental
	result := r.Db.Model(&models.Rental{}).Where("id = ?", id).Preload("Location").Preload("RentalStatus").Preload("RentalRooms").Preload("Photos").Preload("BookingDurationRule").Preload("Bookings").Preload("BookingCostItems").First(&rental)
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
