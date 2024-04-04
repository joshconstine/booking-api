package services

import (
	requests "booking-api/data/request"

	responses "booking-api/data/response"
)

type AmenityTypeService interface {
	FindAll() []responses.AmenityTypeResponse
	FindById(id uint) responses.AmenityTypeResponse
	Create(amenityType requests.CreateAmenityTypeRequest) responses.AmenityTypeResponse
}
