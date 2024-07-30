package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type bookingDetailsRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingDetailsRepositoryImplementation(Db *gorm.DB) BookingDetailsRepository {
	return &bookingDetailsRepositoryImplementation{Db: Db}
}

func (t *bookingDetailsRepositoryImplementation) FindById(id uint) response.BookingDetailsResponse {
	var bookingDetails models.BookingDetails
	result := t.Db.Model(&models.BookingDetails{}).First(&bookingDetails, id)
	if result.Error != nil {
		return response.BookingDetailsResponse{}
	}

	return bookingDetails.MapBookingDetailsToResponse()
}

func (t *bookingDetailsRepositoryImplementation) FindByBookingId(id string) response.BookingDetailsResponse {
	var bookingDetails models.BookingDetails
	result := t.Db.Model(&models.BookingDetails{}).Where("booking_id = ?", id).First(&bookingDetails)
	if result.Error != nil {
		return response.BookingDetailsResponse{}
	}

	return bookingDetails.MapBookingDetailsToResponse()

}

func (t *bookingDetailsRepositoryImplementation) Create(details models.BookingDetails) response.BookingDetailsResponse {
	result := t.Db.Create(&details)
	if result.Error != nil {
		return response.BookingDetailsResponse{}
	}

	return details.MapBookingDetailsToResponse()
}

func (t *bookingDetailsRepositoryImplementation) Update(details request.UpdateBookingDetailsRequest) (response.BookingDetailsResponse, error) {

	var bookingDetails models.BookingDetails

	existingDetails := t.FindById(details.ID)

	bookingDetails.ID = details.ID
	bookingDetails.BookingID = existingDetails.BookingID
	bookingDetails.PaymentComplete = details.PaymentComplete
	bookingDetails.DepositPaid = details.DepositPaid
	bookingDetails.PaymentDueDate = details.PaymentDueDate
	bookingDetails.DocumentsSigned = details.DocumentsSigned
	bookingDetails.BookingStartDate = details.BookingStartDate
	bookingDetails.GuestCount = details.GuestCount

	result := t.Db.Model(&models.BookingDetails{}).Where("id = ?", details.ID).Updates(&bookingDetails)
	if result.Error != nil {
		return response.BookingDetailsResponse{}, result.Error
	}

	return bookingDetails.MapBookingDetailsToResponse(), nil
}
