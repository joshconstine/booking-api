package services

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

func (t BookingServiceImplementation) Create(request requests.CreateUserRequest) (response.BookingResponse, error) {
	// validate request
	err := t.Validate.Struct(request)

	//check if this user already exists
	user := t.UserService.FindByEmail(request.Email)

	bookingToCreate := models.Booking{
		BookingStatus: models.BookingStatus{
			Model: gorm.Model{
				ID: 1,
			},
		},
		Details: models.BookingDetails{
			PaymentComplete:  false,
			DocumentsSigned:  false,
			DepositPaid:      false,
			BookingStartDate: time.Now(),
			PaymentDueDate:   time.Now().AddDate(0, 0, 7),
			LocationID:       1,
		},
		User: models.User{
			Model: gorm.Model{
				ID: 0,
			},
			Email: request.Email,
		},
	}

	if user.Email != request.Email {
		//if not create a new user
		createdUser := t.UserService.CreateUser(request)

		bookingToCreate.UserID = createdUser.ID
	} else {
		bookingToCreate.UserID = user.ID
	}
	if err != nil {
		return response.BookingResponse{}, err
	}
	// create booking
	booking := t.BookingRepository.Create(bookingToCreate)

	bookingToCreate.Details.BookingID = bookingToCreate.ID

	booking = t.BookingRepository.Update(bookingToCreate)

	bookingResponse := booking.MapBookingToResponse()

	return bookingResponse, nil
}
