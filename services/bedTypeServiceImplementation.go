package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BedTypeServiceImplementation struct {
	bedTypeRepository repositories.BedTypeRepository
	Validate          *validator.Validate
}

func NewBedTypeServiceImplementation(bedTypeRepository repositories.BedTypeRepository, validate *validator.Validate) BedTypeService {
	return &BedTypeServiceImplementation{
		bedTypeRepository: bedTypeRepository,
		Validate:          validate,
	}
}

func (t BedTypeServiceImplementation) FindAll() []response.BedTypeResponse {
	result := t.bedTypeRepository.FindAll()

	return result
}

func (t BedTypeServiceImplementation) FindById(bedTypeId uint) response.BedTypeResponse {

	bedTypeData := t.bedTypeRepository.FindById(bedTypeId)

	return bedTypeData
}
