package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type bookingDetailsRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingDetailsRepositoryImplementation(Db *gorm.DB) BookingDetailsRepository {
	return &bookingDetailsRepositoryImplementation{Db: Db}
}

func (t *bookingDetailsRepositoryImplementation) FindById(id uint) models.BookingDetails {
	var bookingDetails models.BookingDetails
	result := t.Db.Where("id = ?", id).First(&bookingDetails)
	if result.Error != nil {
		return models.BookingDetails{}
	}

	return bookingDetails
}
