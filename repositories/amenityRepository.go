package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type AmenityRepository interface {
	FindAll() []response.AmenityResponse
	FindById(id uint) response.AmenityResponse
	Create(amenity requests.CreateAmenityRequest) response.AmenityResponse
}
