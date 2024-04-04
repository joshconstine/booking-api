package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type BedTypeRepository interface {
	FindAll() []response.BedTypeResponse
	FindById(id uint) response.BedTypeResponse
	Create(name models.BedType) response.BedTypeResponse
}
