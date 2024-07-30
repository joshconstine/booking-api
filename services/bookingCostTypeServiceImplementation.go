package services

import (
	responses "booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BookingCostTypeServiceImplementation struct {
	bookingCostTypeRepository repositories.BookingCostTypeRepository
	validator                 *validator.Validate
}

func NewBookingCostTypeServiceImplementation(bookingCostTypeRepository repositories.BookingCostTypeRepository, validator *validator.Validate) BookingCostTypeServiceImplementation {
	return BookingCostTypeServiceImplementation{bookingCostTypeRepository: bookingCostTypeRepository, validator: validator}
}

func (s BookingCostTypeServiceImplementation) FindAll() []responses.BookingCostTypeResponse {
	result := s.bookingCostTypeRepository.FindAll()

	var bookingCostTypeResponses []responses.BookingCostTypeResponse
	for _, bookingCostType := range result {
		bookingCostTypeResponses = append(bookingCostTypeResponses, responses.BookingCostTypeResponse{
			ID:   bookingCostType.ID,
			Name: bookingCostType.Name,
		})
	}
	return bookingCostTypeResponses

}

func (s BookingCostTypeServiceImplementation) FindById(id uint) responses.BookingCostTypeResponse {
	result := s.bookingCostTypeRepository.FindById(id)

	var bookingCostTypeResponse responses.BookingCostTypeResponse
	bookingCostTypeResponse.ID = result.ID
	bookingCostTypeResponse.Name = result.Name
	return bookingCostTypeResponse
}
