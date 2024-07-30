package services

import (
	responses "booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type bookingStatusServiceImplementation struct {
	BookingStatusRepository repositories.BookingStatusRepository
	Validate                *validator.Validate
}

func NewBookingStatusService(bookingStatusRepository repositories.BookingStatusRepository, validate *validator.Validate) BookingStatusService {
	return &bookingStatusServiceImplementation{
		BookingStatusRepository: bookingStatusRepository,
		Validate:                validate,
	}
}

func (t *bookingStatusServiceImplementation) FindAll() []responses.BookingStatusResponse {
	result := t.BookingStatusRepository.FindAll()

	var bookingStatuses []responses.BookingStatusResponse
	for _, value := range result {
		bookingStatus := responses.BookingStatusResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		bookingStatuses = append(bookingStatuses, bookingStatus)
	}
	return bookingStatuses

}

func (t *bookingStatusServiceImplementation) FindById(id uint) responses.BookingStatusResponse {
	result := t.BookingStatusRepository.FindById(id)

	bookingStatus := responses.BookingStatusResponse{
		ID:   result.ID,
		Name: result.Name,
	}

	return bookingStatus
}
