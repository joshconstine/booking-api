package repositories

import (
	"booking-api/models"
)

type BoatRepository interface {
	FindAll() []models.Boat
	FindByID(id int) models.Boat
}
