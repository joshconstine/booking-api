package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	responses "booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type AmenityRepositoryImplementation struct {
	Db *gorm.DB
}

func NewAmenityRepositoryImplementation(db *gorm.DB) AmenityRepository {

	return &AmenityRepositoryImplementation{Db: db}

}

func (t *AmenityRepositoryImplementation) FindAll() []responses.AmenityResponse {
	var amenities []models.Amenity
	var response []response.AmenityResponse

	result := t.Db.Preload("AmenityType").Find(&amenities)
	if result.Error != nil {

	}

	var item responses.AmenityResponse
	for _, amenity := range amenities {
		item.ID = amenity.ID
		item.Name = amenity.Name
		item.AmenityType = responses.AmenityTypeResponse{
			ID:   amenity.AmenityType.ID,
			Name: amenity.AmenityType.Name,
		}

		response = append(response, item)

	}
	return response
}

func (t *AmenityRepositoryImplementation) FindById(id uint) response.AmenityResponse {
	var amenity models.Amenity
	var response response.AmenityResponse

	result := t.Db.Preload("AmenityType").Find(&amenity)
	if result.Error != nil {

		return response
	}

	response.ID = amenity.ID
	response.Name = amenity.Name
	response.AmenityType = responses.AmenityTypeResponse{
		ID:   amenity.AmenityType.ID,
		Name: amenity.AmenityType.Name,
	}

	return response
}

func (t *AmenityRepositoryImplementation) Create(amenity requests.CreateAmenityRequest) response.AmenityResponse {

	amenityToCreate := models.Amenity{
		Name:        amenity.Name,
		AmenityType: models.AmenityType{Model: gorm.Model{ID: amenity.AmenityTypeId}},
	}

	result := t.Db.Table("amenities").Create(&amenityToCreate)
	if result.Error != nil {
		return response.AmenityResponse{}
	}

	return response.AmenityResponse{

		Name: amenity.Name,
		AmenityType: response.AmenityTypeResponse{
			ID: amenity.AmenityTypeId,
		},
	}
}
