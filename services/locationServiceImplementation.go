package services

import (
	responses "booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type LocationServiceImplementation struct {
	LocationRepository repositories.LocationRepository
	Validate           *validator.Validate
}

func NewLocationServiceImplementation(locationRepository repositories.LocationRepository, validate *validator.Validate) LocationService {
	return &LocationServiceImplementation{
		LocationRepository: locationRepository,
		Validate:           validate,
	}
}

func (t LocationServiceImplementation) FindAll() []responses.LocationResponse {
	result := t.LocationRepository.FindAll()

	return result
}

func (t LocationServiceImplementation) FindById(id uint) responses.LocationResponse {
	result := t.LocationRepository.FindById(id)

	return result
}

func (t LocationServiceImplementation) Create(locationName string) responses.LocationResponse {
	location := t.LocationRepository.Create(locationName)

	return location
}
