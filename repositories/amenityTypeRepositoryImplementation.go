package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type AmenityTypeRepositoryImplementation struct {
	Db *gorm.DB
}

func NewAmenityTypeRepositoryImplementation(db *gorm.DB) AmenityTypeRepository {
	return &AmenityTypeRepositoryImplementation{Db: db}
}

func (t *AmenityTypeRepositoryImplementation) FindAll() []models.AmenityType {
	var amenityTypes []models.AmenityType
	result := t.Db.Find(&amenityTypes)
	if result.Error != nil {
		return []models.AmenityType{}
	}

	return amenityTypes
}

func (t *AmenityTypeRepositoryImplementation) FindById(id uint) models.AmenityType {
	var amenityType models.AmenityType
	result := t.Db.Where("id = ?", id).First(&amenityType)
	if result.Error != nil {
		return models.AmenityType{}
	}

	return amenityType
}

func (t *AmenityTypeRepositoryImplementation) Create(amenityType models.AmenityType) models.AmenityType {
	result := t.Db.Create(&amenityType)
	if result.Error != nil {
		return models.AmenityType{}
	}

	return amenityType
}
