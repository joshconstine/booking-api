package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/repositories"
)

type BookingDetailsServiceImplementation struct {
	bookingDetailsRepository repositories.BookingDetailsRepository
	bookingPaymentRepository repositories.BookingPaymentRepository
	bookingService           BookingService
}

func NewBookingDetailsServiceImplementation(bookingDetailsRepository repositories.BookingDetailsRepository, bookingPaymentRepository repositories.BookingPaymentRepository, bookingService BookingService) *BookingDetailsServiceImplementation {
	return &BookingDetailsServiceImplementation{bookingDetailsRepository: bookingDetailsRepository, bookingPaymentRepository: bookingPaymentRepository, bookingService: bookingService}
}

func (service BookingDetailsServiceImplementation) FindById(id uint) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.FindById(id)
}
func (service BookingDetailsServiceImplementation) FindByBookingId(id string) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.FindByBookingId(id)
}

func (service BookingDetailsServiceImplementation) Create(details models.BookingDetails) response.BookingDetailsResponse {
	return service.bookingDetailsRepository.Create(details)
}

func (service BookingDetailsServiceImplementation) Update(details request.UpdateBookingDetailsRequest) (response.BookingDetailsResponse, error) {
	return service.bookingDetailsRepository.Update(details)
}

func (service BookingDetailsServiceImplementation) AuditBookingDetailsForBooking(bookingId string) error {
	//Get booking information
	booking, err := service.bookingService.FindById(bookingId)
	if err != nil {
		return err
	}

	//Audit PaymentStatus
	outstandingAmount := service.bookingPaymentRepository.FindTotalOutstandingAmountByBookingId(bookingId)
	if outstandingAmount == 0 {
		//ensure booking status is paid
		//update booking status to paid
		if booking.Details.PaymentComplete == false {
			booking.Details.PaymentComplete = true
			_, err := service.Update(request.UpdateBookingDetailsRequest{
				ID:               booking.Details.ID,
				PaymentComplete:  true,
				BookingStartDate: booking.Details.BookingStartDate,
				PaymentDueDate:   booking.Details.PaymentDueDate,
				DocumentsSigned:  booking.Details.DocumentsSigned,
				DepositPaid:      true,
				GuestCount:       booking.Details.GuestCount,
			})
			if err != nil {
				return err
			}

		}
	}
	//Audit BookingStatus

	return nil
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

func AuditBookingStatus(bookingInformation response.BookingInformationResponse) {
	//check if booking is complete
	if bookingInformation.Details.PaymentComplete == true && bookingInformation.Details.DocumentsSigned == true {
		//Check if the booking is in Progress
		//update booking status to in progress

		//update booking status to confirmed

	} else if bookingInformation.Details.PaymentComplete == true {

		//update booking status to payment confirmed
	} else {
		//update booking status to pending
	}
}
