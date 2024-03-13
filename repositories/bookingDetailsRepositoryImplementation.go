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

func (t *bookingDetailsRepositoryImplementation) Create(details models.BookingDetails) models.BookingDetails {
	result := t.Db.Create(&details)
	if result.Error != nil {
		return models.BookingDetails{}
	}

	return details
}

func (t *bookingDetailsRepositoryImplementation) Update(details models.BookingDetails) models.BookingDetails {
	result := t.Db.Save(&details)
	if result.Error != nil {
		return models.BookingDetails{}
	}

	return details
}
