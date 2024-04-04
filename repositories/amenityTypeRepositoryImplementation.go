package repositories

import (
	"booking-api/data/response"
	responses "booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type AmenityTypeRepositoryImplementation struct {
	Db *gorm.DB
}

func NewAmenityTypeRepositoryImplementation(db *gorm.DB) AmenityTypeRepository {
	return &AmenityTypeRepositoryImplementation{Db: db}
}

func (t *AmenityTypeRepositoryImplementation) FindAll() []responses.AmenityTypeResponse {
	var amenityTypes []models.AmenityType
	result := t.Db.Find(&amenityTypes)
	if result.Error != nil {
		return []responses.AmenityTypeResponse{}
	}

	var response []responses.AmenityTypeResponse
	var item responses.AmenityTypeResponse
	for _, amenityType := range amenityTypes {
		item.ID = amenityType.ID
		item.Name = amenityType.Name

		response = append(response, item)
	}
	return response
}

func (t *AmenityTypeRepositoryImplementation) FindById(id uint) responses.AmenityTypeResponse {
	var amenityType models.AmenityType
	result := t.Db.Where("id = ?", id).First(&amenityType)
	if result.Error != nil {
		return response.AmenityTypeResponse{}
	}

	var response = response.AmenityTypeResponse{
		ID:   amenityType.ID,
		Name: amenityType.Name,
	}

	return response
}

func (t *AmenityTypeRepositoryImplementation) Create(amenityType models.AmenityType) responses.AmenityTypeResponse {
	result := t.Db.Create(&amenityType)
	if result.Error != nil {
		return responses.AmenityTypeResponse{}
	}

	var response = responses.AmenityTypeResponse{
		ID:   amenityType.ID,
		Name: amenityType.Name,
	}

	return response

}
