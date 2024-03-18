package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type LocationRepositoryImplementation struct {
	Db *gorm.DB
}

func NewLocationRepositoryImplementation(db *gorm.DB) LocationRepository {
	return &LocationRepositoryImplementation{Db: db}
}

func (t *LocationRepositoryImplementation) FindAll() []models.Location {
	var locations []models.Location
	result := t.Db.Find(&locations)
	if result.Error != nil {
		return []models.Location{}
	}

	return locations
}

func (t *LocationRepositoryImplementation) FindById(id uint) models.Location {
	var location models.Location
	result := t.Db.Where("id = ?", id).First(&location)
	if result.Error != nil {
		return models.Location{}
	}

	return location
}

func (t *LocationRepositoryImplementation) Create(location models.Location) models.Location {

	result := t.Db.Create(&location)
	if result.Error != nil {
		return models.Location{}
	}

	return location
}
