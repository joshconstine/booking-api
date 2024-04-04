package repositories

import (
	"booking-api/models"
)

type BoatRepository interface {
	FindAll() []models.Boat
	FindById(id int) models.Boat
}