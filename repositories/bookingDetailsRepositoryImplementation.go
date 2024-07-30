package repositories

import (
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
	result := t.Db.First(&bookingDetails, id)
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

func (t *bookingDetailsRepositoryImplementation) Update(details models.BookingDetails) response.BookingDetailsResponse {
	result := t.Db.Save(&details)
	if result.Error != nil {
		return response.BookingDetailsResponse{}
	}

	return details.MapBookingDetailsToResponse()
}
