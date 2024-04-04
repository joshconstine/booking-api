package services

import (
	responses "booking-api/data/response"
)

type BedTypeService interface {
	FindAll() []responses.BedTypeResponse
	FindById(id uint) responses.BedTypeResponse
}
