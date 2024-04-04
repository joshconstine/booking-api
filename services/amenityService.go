package services

import (
	requests "booking-api/data/request"

	responses "booking-api/data/response"
)

type AmenityService interface {
	FindAll() []responses.AmenityResponse
	FindById(id uint) responses.AmenityResponse
	Create(amenity requests.CreateAmenityRequest) responses.AmenityResponse
}
