package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BookingServiceImplementation struct {
	BookingRepository repositories.BookingRepository
	UserService       UserService
	Validate          *validator.Validate
}

func NewBookingServiceImplementation(bookingRepository repositories.BookingRepository, validate *validator.Validate, userService UserService) BookingService {
	return &BookingServiceImplementation{
		BookingRepository: bookingRepository,
		Validate:          validate,
		UserService:       userService,
	}
}

func (t BookingServiceImplementation) FindAll() []response.BookingResponse {
	result := t.BookingRepository.FindAll()
	return result
}

func (t BookingServiceImplementation) FindById(id string) response.BookingInformationResponse {
	result := t.BookingRepository.FindById(id)

	return result
}

func (t BookingServiceImplementation) Create(request *request.CreateBookingRequest) error {

	bookingToCreate := request.MapCreateBookingRequestToBooking()

	t.BookingRepository.Create(&bookingToCreate)

	return nil
}
