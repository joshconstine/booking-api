package repositories

import (
	"booking-api/data/response"
)

type LocationRepository interface {
	FindAll() []response.LocationResponse
	FindById(id uint) response.LocationResponse
	Create(locationName string) response.LocationResponse
}
