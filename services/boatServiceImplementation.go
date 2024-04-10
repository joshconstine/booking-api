package services

import (
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BoatServiceImplementation struct {
	BoatRepository repositories.BoatRepository
	Validate       *validator.Validate
}

func NewBoatServiceImplementation(boatRepository repositories.BoatRepository, validate *validator.Validate) BoatService {
	return &BoatServiceImplementation{
		BoatRepository: boatRepository,
		Validate:       validate,
	}
}

func (t BoatServiceImplementation) FindAll() []response.BoatResponse {
	result := t.BoatRepository.FindAll()

	return result
}

func (t BoatServiceImplementation) FindById(id int) response.BoatInformationResponse {
	result := t.BoatRepository.FindById(id)

	return result
}
