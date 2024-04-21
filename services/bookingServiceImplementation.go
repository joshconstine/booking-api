package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
	"fmt"

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

	//check if the entities are available during the timeblocks
	//if the entity requires does not allow instant booking, check if there was in inquery > entityBookingRequest was approved.
	//if it was approved, the request.entitybookingrequest  start and end times must be within the approved time ^
	//Verify the user exists
	//Add all booking cost items, documents
	//set booking payment due data
	//send email to user
	//send email to account owner

	canBook, err := t.BookingRepository.CheckIfEntitiesCanBeBooked(request)

	if err != nil {
		fmt.Printf("error checking if entities can be booked: %v", err)
		return err
	}

	if !canBook {
		return err
	} else {

		bookingToCreate := request.MapCreateBookingRequestToBooking()

		t.BookingRepository.Create(&bookingToCreate)
	}

	return nil
}
