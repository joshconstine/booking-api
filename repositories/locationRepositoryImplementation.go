package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type LocationRepositoryImplementation struct {
	Db *gorm.DB
}

func NewLocationRepositoryImplementation(db *gorm.DB) LocationRepository {
	return &LocationRepositoryImplementation{Db: db}
}

func (t *LocationRepositoryImplementation) FindAll() []response.LocationResponse {
	var locations []models.Location
	result := t.Db.Find(&locations)
	if result.Error != nil {
		return []response.LocationResponse{}
	}

	var locationResponses []response.LocationResponse
	for _, location := range locations {
		locationResponses = append(locationResponses, response.LocationResponse{
			ID:   location.ID,
			Name: location.Name,
		})
	}

	return locationResponses
}

func (t *LocationRepositoryImplementation) FindById(id uint) response.LocationResponse {
	var location models.Location
	result := t.Db.Where("id = ?", id).First(&location)
	if result.Error != nil {
		return response.LocationResponse{}
	}

	return response.LocationResponse{
		ID:   location.ID,
		Name: location.Name,
	}

}

func (t *LocationRepositoryImplementation) Create(locationName string) response.LocationResponse {

	location := models.Location{
		Name: locationName,
	}

	result := t.Db.Create(&location)
	if result.Error != nil {
		return response.LocationResponse{}
	}

	return response.LocationResponse{
		ID:   location.ID,
		Name: location.Name,
	}
}
