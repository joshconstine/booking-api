package services

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/pkg/email"
	"booking-api/repositories"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type BookingServiceImplementation struct {
	BookingRepository        repositories.BookingRepository
	BookingDetailsRepository repositories.BookingDetailsRepository
	BookingPaymentRepository repositories.BookingPaymentRepository
	UserService              UserService
	EntityBookingService     EntityBookingService
	Validate                 *validator.Validate
}

func NewBookingServiceImplementation(bookingRepository repositories.BookingRepository, bookingDetailsRepository repositories.BookingDetailsRepository, bookingPaymentRepository repositories.BookingPaymentRepository, validate *validator.Validate, userService UserService, entityBookingService EntityBookingService) BookingService {
	return &BookingServiceImplementation{
		BookingRepository:        bookingRepository,
		BookingDetailsRepository: bookingDetailsRepository,
		BookingPaymentRepository: bookingPaymentRepository,
		Validate:                 validate,
		UserService:              userService,
		EntityBookingService:     entityBookingService,
	}
}

func (t BookingServiceImplementation) FindAll() []response.BookingResponse {
	result := t.BookingRepository.FindAll()
	return result
}

func (t BookingServiceImplementation) FindById(id string) (response.BookingInformationResponse, error) {
	result := t.BookingRepository.FindById(id)

	customerInfo, err := t.UserService.FindByPublicUserID(result.Customer.UserID)

	if err != nil {
		fmt.Printf("error finding customer by public user id: %v", err)
		return response.BookingInformationResponse{}, err
	}

	result.Customer = customerInfo

	return result, nil
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

	user, err := t.UserService.FindByEmail(createBookingRequest.Email)
	if err != nil {
		return "", err

	}

	createBookingRequest.UserID = user.UserID
	createBookingRequest.FirstName = user.FirstName
	createBookingRequest.PhoneNumber = user.PhoneNumber

	//validate the request
	err = t.Validate.Struct(createBookingRequest)
	if err != nil {
		return "", err

	}

	return t.Create(&createBookingRequest)
}

func (t BookingServiceImplementation) GetSnapshot(request request.GetBookingSnapshotRequest) []response.BookingSnapshotResponse {
	result := t.BookingRepository.GetSnapshot(request)
	return result
}

func (t BookingServiceImplementation) UpdateBookingStatusForBooking(request request.UpdateBookingStatusRequest) error {
	return t.BookingRepository.UpdateBookingStatusForBooking(request)
}

//+-----------+
//|name       |
//+-----------+
//|Drafted    |
//|Requested  |
//|Confirmed  |
//|In Progress|
//|Completed  |
//|Cancelled  |
//+-----------+

func isAtLeastOneEntityBookingInProgress(entitiyBookings []response.EntityBookingResponse) bool {
	result := false
	for _, entityBooking := range entitiyBookings {
		if entityBooking.Status.ID == constants.BOOKING_STATUS_IN_PROGRESS_ID {
			result = true
		}

	}
	return result
}
func areAllEntityBookingsCompleted(entitiyBookings []response.EntityBookingResponse) bool {
	result := true
	for _, entityBooking := range entitiyBookings {
		if entityBooking.Status.ID != constants.BOOKING_STATUS_COMPLETED_ID {
			result = false
		}

	}
	return result
}

func (t BookingServiceImplementation) AuditBookingStatusForBooking(bookingInformation response.BookingInformationResponse) {
	//check if booking is complete
	var request request.UpdateBookingStatusRequest
	request.BookingID = bookingInformation.ID

	//check if should be drafted
	//if bookingInformation.Status.ID == constants.BOOKING_STATUS_COMPLETED_ID || bookingInformation.Status.ID == constants.BOOKING_STATUS_CANCELLED_ID || len(bookingInformation.Entities) == 0 {
	if len(bookingInformation.Entities) == 0 {
		return

	}
	//audit payment status
	t.AuditPaymentStatusForBooking(&bookingInformation)

	//audit document signed status
	t.AuditDocumentSignedStatusForBooking(&bookingInformation)

	//audit each entity booking status
	for _, entityBooking := range bookingInformation.Entities {
		t.EntityBookingService.AuditEntityBookingStatusForBooking(&bookingInformation, &entityBooking)

	}

	if bookingInformation.Details.PaymentComplete == true && bookingInformation.Details.DocumentsSigned == true && bookingInformation.Status.ID != constants.BOOKING_STATUS_COMPLETED_ID {
		//Check if the booking is in Progress
		//update booking status to in progress
		if isAtLeastOneEntityBookingInProgress(bookingInformation.Entities) {
			request.BookingStatusID = constants.BOOKING_STATUS_IN_PROGRESS_ID
			if bookingInformation.Status.ID != constants.BOOKING_STATUS_IN_PROGRESS_ID {
				t.UpdateBookingStatusForBooking(request)
			}
		} else {

			request.BookingStatusID = constants.BOOKING_STATUS_CONFIRMED_ID
			if bookingInformation.Status.ID != constants.BOOKING_STATUS_CONFIRMED_ID {
				t.UpdateBookingStatusForBooking(request)
			}
		}

		//update booking status to confirmed

	} else if bookingInformation.Details.PaymentComplete == true && bookingInformation.Details.DocumentsSigned == true && bookingIsInPast(bookingInformation) {
		if bookingInformation.Status.ID != constants.BOOKING_STATUS_COMPLETED_ID {
			request.BookingStatusID = constants.BOOKING_STATUS_COMPLETED_ID
			t.UpdateBookingStatusForBooking(request)
		}

	}

	if areAllEntityBookingsCompleted(bookingInformation.Entities) {
		request.BookingStatusID = constants.BOOKING_STATUS_COMPLETED_ID
		if bookingInformation.Status.ID != constants.BOOKING_STATUS_COMPLETED_ID {
			t.UpdateBookingStatusForBooking(request)
		}

	}
}
func bookingIsInPast(bookingInformation response.BookingInformationResponse) bool {
	var today = time.Now()
	isInPast := true

	for _, entity := range bookingInformation.Entities {
		if entity.Timeblock.EndTime.After(today) {
			isInPast = false
		}

	}
	return isInPast
}
func (t BookingServiceImplementation) AuditDocumentSignedStatusForBooking(bookingInformation *response.BookingInformationResponse) {
	//check if booking is complete
	var request request.UpdateBookingDocumentsSignedRequest
	request.BookingID = bookingInformation.ID

	var documentsSigned bool
	documentsSigned = true

	for _, document := range bookingInformation.Documents {
		if document.Signed == false && document.RequiresSignature == true {
			documentsSigned = false
		}
	}

	if documentsSigned != bookingInformation.Details.DocumentsSigned {
		request.DocumentsSigned = documentsSigned
		bookingInformation.Details.DocumentsSigned = documentsSigned
		t.BookingRepository.UpdateBookingDocumentsSigned(request)
	}

}
func (t BookingServiceImplementation) AuditPaymentStatusForBooking(booking *response.BookingInformationResponse) {

	//check if there are cost items
	if len(booking.CostItems) == 0 {
		if booking.Details.PaymentComplete == true || booking.Details.DepositPaid == true {
			_, err := t.BookingDetailsRepository.Update(request.UpdateBookingDetailsRequest{
				ID:               booking.Details.ID,
				PaymentComplete:  false,
				BookingStartDate: booking.Details.BookingStartDate,
				PaymentDueDate:   booking.Details.PaymentDueDate,
				DocumentsSigned:  booking.Details.DocumentsSigned,
				DepositPaid:      false,
				GuestCount:       booking.Details.GuestCount,
			})
			if err != nil {
				//return err
			}
		}
		return
	}

	//Audit PaymentStatus

	outstandingAmount := t.BookingPaymentRepository.FindTotalOutstandingAmountByBookingId(booking.ID)

	if outstandingAmount == 0 {
		//ensure booking status is paid
		//update booking status to paid
		if booking.Details.PaymentComplete == false {
			booking.Details.PaymentComplete = true
			_, err := t.BookingDetailsRepository.Update(request.UpdateBookingDetailsRequest{
				ID:               booking.Details.ID,
				PaymentComplete:  true,
				BookingStartDate: booking.Details.BookingStartDate,
				PaymentDueDate:   booking.Details.PaymentDueDate,
				DocumentsSigned:  booking.Details.DocumentsSigned,
				DepositPaid:      true,
				GuestCount:       booking.Details.GuestCount,
			})
			if err != nil {
				//return err
			}

		}
	}
}
func (t BookingServiceImplementation) AuditAllBookingStatus() error {
	bookings := t.FindAll()
	var bookingInformation response.BookingInformationResponse
	for _, booking := range bookings {
		bookingInformation, _ = t.FindById(booking.ID)
		t.AuditBookingStatusForBooking(bookingInformation)

	}
	return nil
}
