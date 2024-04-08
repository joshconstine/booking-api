package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type BoatRepository interface {
	FindAll() []response.BoatResponse
	FindById(id int) response.BoatResponse
	Create(boat models.Boat) models.Boat
}
