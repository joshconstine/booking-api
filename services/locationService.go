package services

import (
	"booking-api/data/response"
)

type LocationService interface {
	FindAll() []response.LocationResponse
	FindById(id uint) response.LocationResponse
	Create(locationName string) response.LocationResponse
}
