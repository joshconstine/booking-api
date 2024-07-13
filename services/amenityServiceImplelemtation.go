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
func (t AmenityServiceImplementation) FindAllSorted() []response.SortedAmenityResponse {

	allAmenities := t.AmenityRepository.FindAll()

	var sortedAmenities []response.SortedAmenityResponse

	for _, amenity := range allAmenities {
		if len(sortedAmenities) == 0 {
			sortedAmenities = append(sortedAmenities, response.SortedAmenityResponse{
				TypeId:    amenity.AmenityType.ID,
				TypeName:  amenity.AmenityType.Name,
				Amenities: []response.AmenityResponse{amenity},
			})
		} else {
			for index, sortedAmenity := range sortedAmenities {
				if sortedAmenity.TypeId == amenity.AmenityType.ID {
					sortedAmenities[index].Amenities = append(sortedAmenities[index].Amenities, amenity)
					break
				} else if index == len(sortedAmenities)-1 {
					sortedAmenities = append(sortedAmenities, response.SortedAmenityResponse{
						TypeId:    amenity.AmenityType.ID,
						TypeName:  amenity.AmenityType.Name,
						Amenities: []response.AmenityResponse{amenity},
					})
					break
				}

			}
		}
	}

	return sortedAmenities
}

func (t AmenityServiceImplementation) FindById(amenityId uint) response.AmenityResponse {
	amenityData := t.AmenityRepository.FindById(amenityId)

	amenityResponse := response.AmenityResponse{
		ID:   amenityData.ID,
		Name: amenityData.Name,
	}
	return amenityResponse
}
