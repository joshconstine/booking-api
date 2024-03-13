package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
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

	var bookings []response.BookingResponse
	for _, value := range result {
		booking := response.BookingResponse{
			ID:               value.ID,
			BookingStatusID:  value.BookingStatusID,
			BookingDetailsID: value.BookingDetailsID,
		}
		bookings = append(bookings, booking)
	}
	return bookings
}

func (t BookingServiceImplementation) FindById(id uint) response.BookingResponse {
	result := t.BookingRepository.FindById(id)

	booking := response.BookingResponse{
		ID:               result.ID,
		BookingStatusID:  result.BookingStatusID,
		BookingDetailsID: result.BookingDetailsID,
	}
	return booking
}

func (t BookingServiceImplementation) Create(request requests.CreateUserRequest) response.BookingResponse {
	// validate request
	err := t.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	var bookingToCreate = models.Booking{}

	//

	//check if this user already exists
	user := t.UserService.FindByEmail(request.Email)
	if user.Email != request.Email {
		//if not create a new user
		createdUser := t.UserService.CreateUser(request)

		// request.ID = int(userID)
		bookingToCreate.UserID = createdUser.ID
	} else {
		// request.ID = user.ID
		bookingToCreate.UserID = user.ID
	}

	// create booking
	booking := t.BookingRepository.Create(bookingToCreate)

	// return response
	bookingResponse := response.BookingResponse{
		ID:               booking.ID,
		BookingStatusID:  booking.BookingStatusID,
		BookingDetailsID: booking.BookingDetailsID,
	}
	return bookingResponse
}
