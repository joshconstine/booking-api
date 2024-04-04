package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type AmenityServiceImplementation struct {
	AmenityRepository repositories.AmenityRepository
	Validate          *validator.Validate
}

func NewAmenityServiceImplementation(amenityRepository repositories.AmenityRepository, validate *validator.Validate) AmenityService {
	return &AmenityServiceImplementation{
		AmenityRepository: amenityRepository,
		Validate:          validate,
	}
}

func (t AmenityServiceImplementation) Create(amenity requests.CreateAmenityRequest) response.AmenityResponse {
	err := t.Validate.Struct(amenity)

	if err != nil {
		panic(err)
	}

	return t.AmenityRepository.Create(amenity)

}

func (t AmenityServiceImplementation) FindAll() []response.AmenityResponse {
	result := t.AmenityRepository.FindAll()

	var amenities []response.AmenityResponse
	for _, value := range result {
		amenity := response.AmenityResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		amenities = append(amenities, amenity)
	}
	return amenities
}

func (t AmenityServiceImplementation) FindById(amenityId uint) response.AmenityResponse {
	amenityData := t.AmenityRepository.FindById(amenityId)

	amenityResponse := response.AmenityResponse{
		ID:   amenityData.ID,
		Name: amenityData.Name,
	}
	return amenityResponse
}
