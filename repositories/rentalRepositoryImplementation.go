package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type RentalRepositoryImplementation struct {
	Db *gorm.DB
}

func NewRentalRepositoryImplementation(db *gorm.DB) *RentalRepositoryImplementation {
	return &RentalRepositoryImplementation{Db: db}
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
