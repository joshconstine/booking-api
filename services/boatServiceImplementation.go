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

// func (t BoatServiceImplementation) Create(tag request.CreateBoatsRequest) {
// 	err := t.Validate.Struct(tag)
// 	helper.ErrorPanic(err)
// 	tagModel := model.Boats{
// 		Name: tag.Name,
// 	}
// 	t.BoatRepository.Save(tagModel)
// }

// func (t BoatServiceImplementation) Update(tag request.UpdateBoatsRequest) {
// 	tagData, err := t.BoatRepository.FindById(tag.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = tag.Name
// 	t.BoatRepository.Update(tagData)
// }

// func (t BoatServiceImplementation) Delete(tagId int) {
// 	t.BoatRepository.Delete(tagId)
// }

// func (t BoatServiceImplementation) FindById(tagId int) response.BoatsResponse {
// 	tagData, err := t.BoatRepository.FindById(tagId)
// 	helper.ErrorPanic(err)

// 	tagResponse := response.BoatsResponse{
// 		Id:   tagData.Id,
// 		Name: tagData.Name,
// 	}
// 	return tagResponse
// }

func (t BoatServiceImplementation) FindAll() []response.BoatResponse {
	result := t.BoatRepository.FindAll()

	var boats []response.BoatResponse
	for _, value := range result {
		boat := response.BoatResponse{
			ID:        value.ID,
			Name:      value.Name,
			Occupancy: value.Occupancy,
			MaxWeight: value.MaxWeight,
			Photos:    nil,
		}
		boats = append(boats, boat)
	}
	return boats
}

func (t BoatServiceImplementation) FindById(id int) response.BoatResponse {
	result := t.BoatRepository.FindById(id)

	boat := response.BoatResponse{
		ID:        result.ID,
		Name:      result.Name,
		Occupancy: result.Occupancy,
		MaxWeight: result.MaxWeight,
		Photos:    nil,
	}
	return boat
}
