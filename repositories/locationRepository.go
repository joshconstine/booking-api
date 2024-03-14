package repositories

import (
	"booking-api/models"
)

type LocationRepository interface {
	FindAll() []models.Location
	FindById(id uint) models.Location
	Create(location models.Location) models.Location
}
