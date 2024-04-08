package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BookingServiceImplementation struct {
	BookingRepository     repositories.BookingRepository
	UserService           UserService
	BookingDetailsService BookingDetailsService
	Validate              *validator.Validate
}

func NewBookingServiceImplementation(bookingRepository repositories.BookingRepository, validate *validator.Validate, userService UserService, bookingDetailsService BookingDetailsService) BookingService {
	return &BookingServiceImplementation{
		BookingRepository:     bookingRepository,
		Validate:              validate,
		UserService:           userService,
		BookingDetailsService: bookingDetailsService,
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

		bookingToCreate.UserID = createdUser.ID
	} else {
		bookingToCreate.UserID = user.ID
	}
	bookingToCreate.User = models.User{
		Email: user.Email,
	}

	bookingToCreate.User = models.User{
		Model: gorm.Model{
			ID: bookingToCreate.UserID,
		},
		Email: request.Email,
	}

	// create booking
	booking := t.BookingRepository.Create(bookingToCreate)

	bookingResponse := booking.MapBookingToResponse()

	return bookingResponse
}
