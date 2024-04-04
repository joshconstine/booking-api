package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type AmenityTypeServiceImplementation struct {
	amenityTypeRepository repositories.AmenityTypeRepository
	Validate              *validator.Validate
}

func NewAmenityTypeServiceImplementation(amenityTypeRepository repositories.AmenityTypeRepository, validate *validator.Validate) AmenityTypeService {
	return &AmenityTypeServiceImplementation{
		amenityTypeRepository: amenityTypeRepository,
		Validate:              validate,
	}
}

func (t AmenityTypeServiceImplementation) FindAll() []response.AmenityTypeResponse {
	result := t.amenityTypeRepository.FindAll()

	return result
}

func (t AmenityTypeServiceImplementation) FindById(amenityTypeId uint) response.AmenityTypeResponse {

	amenityTypeData := t.amenityTypeRepository.FindById(amenityTypeId)

	return amenityTypeData
}

func (t AmenityTypeServiceImplementation) Create(amenityType requests.CreateAmenityTypeRequest) response.AmenityTypeResponse {
	err := t.Validate.Struct(amenityType)

	if err != nil {
		panic(err)
	}

	var amenityTypeToCreate = models.AmenityType{
		Name: amenityType.Name,
	}

	return t.amenityTypeRepository.Create(amenityTypeToCreate)

}
