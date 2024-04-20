package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type BoatRepository interface {
	FindAll() []response.BoatResponse
	FindAllIDs() []uint
	FindById(id int) response.BoatInformationResponse
	Create(boat models.Boat) models.Boat
}
