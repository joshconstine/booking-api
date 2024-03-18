package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type AmenityRepositoryImplementation struct {
	Db *gorm.DB
}

func NewAmenityRepositoryImplementation(db *gorm.DB) AmenityRepository {
	return &AmenityRepositoryImplementation{Db: db}
}

func (t *AmenityRepositoryImplementation) FindAll() []models.Amenity {
	var amenities []models.Amenity
	result := t.Db.Find(&amenities)
	if result.Error != nil {
		return []models.Amenity{}
	}

	return amenities
}

func (t *AmenityRepositoryImplementation) FindById(id uint) models.Amenity {
	var amenity models.Amenity
	result := t.Db.Where("id = ?", id).First(&amenity)
	if result.Error != nil {
		return models.Amenity{}
	}

	return amenity
}

func (t *AmenityRepositoryImplementation) Create(amenity models.Amenity) models.Amenity {
	result := t.Db.Create(&amenity)
	if result.Error != nil {
		return models.Amenity{}
	}

	return amenity
}
