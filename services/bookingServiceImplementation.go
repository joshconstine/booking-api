package services

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/pkg/email"
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

func (t BookingServiceImplementation) Create(request *request.CreateBookingRequest) (string, error) {

	//check if the entities are available during the timeblocks
	//if the entity requires does not allow instant booking, check if there was in inquery > entityBookingRequest was approved.
	//if it was approved, the request.entitybookingrequest  start and end times must be within the approved time ^
	//Verify the user exists
	//Add all booking cost items, documents
	//set booking payment due data
	//send email to user
	//send email to account owner

	canBook, err := t.BookingRepository.CheckIfEntitiesCanBeBooked(request)
	var bookingId string
	if err != nil {
		fmt.Printf("error checking if entities can be booked: %v", err)
		return "", err
	}

	if !canBook {
		return "", err
	} else {
		//Start transaction
		bookingId, err = t.BookingRepository.Create(request)
		if err != nil {
			fmt.Printf("error creating booking: %v", err)
			return "", err
		}

		//End transaction
		//send confirmation email to user
		//check if mode if Production or Development
		email.SendEmailTemplate(constants.APPLICATION_NAME, constants.SEND_GRID_EMAIL, request.FirstName, "joshua.constine97@gmail.com", "Booking Confirmation"+bookingId, constants.EMAIL_TEMPLATE_BOOKING_CONFIRMATION, map[string]interface{}{
			"bookingId": bookingId,
		})

	}

	return bookingId, nil
}

func (t BookingServiceImplementation) CreateBookingWithUserInformation(request *request.CreateBookingWithUserInformationRequest) (string, error) {

	createBookingRequest := request.MapCreateBookingWithUserInformationRequestToCreateBookingRequest()

	//validate the request
	err := t.Validate.Struct(createBookingRequest)
	if err != nil {
		return "", err

	}

	return t.Create(&createBookingRequest)
}

func (t BookingServiceImplementation) GetSnapshot() []response.BookingSnapshotResponse {
	result := t.BookingRepository.GetSnapshot()
	return result
}
