package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type bookingRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingRepositoryImplementation(Db *gorm.DB) BookingRepository {
	return &bookingRepositoryImplementation{Db: Db}
}

func (t *bookingRepositoryImplementation) FindAll() []models.Booking {
	var bookings []models.Booking
	result := t.Db.Find(&bookings)
	if result.Error != nil {
		return []models.Booking{}
	}

	return bookings
}

func (t *bookingRepositoryImplementation) FindById(id string) models.Booking {
	var booking models.Booking
	result := t.Db.Where("id = ?", id).First(&booking)
	if result.Error != nil {
		return models.Booking{}
	}

	return booking
}

func (t *bookingRepositoryImplementation) Create(booking models.Booking) models.Booking {
	result := t.Db.Create(&booking)
	if result.Error != nil {
		return models.Booking{}
	}

	return booking
}

func (t *bookingRepositoryImplementation) Update(booking models.Booking) models.Booking {
	result := t.Db.Save(&booking)
	if result.Error != nil {
		return models.Booking{}
	}

	return booking
}
