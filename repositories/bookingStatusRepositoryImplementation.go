package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type BookingStatusRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingStatusRepositoryImplementation(db *gorm.DB) BookingStatusRepository {
	return &BookingStatusRepositoryImplementation{Db: db}
}

func (t *BookingStatusRepositoryImplementation) FindAll() []models.BookingStatus {
	var bookingStatuses []models.BookingStatus
	result := t.Db.Find(&bookingStatuses)
	if result.Error != nil {
		return []models.BookingStatus{}
	}

	return bookingStatuses
}

func (t *BookingStatusRepositoryImplementation) FindById(id uint) models.BookingStatus {
	var bookingStatus models.BookingStatus
	result := t.Db.Where("id = ?", id).First(&bookingStatus)
	if result.Error != nil {
		return models.BookingStatus{}
	}

	return bookingStatus
}

func (t *BookingStatusRepositoryImplementation) Create(status models.BookingStatus) models.BookingStatus {
	result := t.Db.Create(&status)
	if result.Error != nil {
		return models.BookingStatus{}
	}

	return status
}
