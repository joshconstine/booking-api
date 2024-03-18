package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type BookingCostTypeRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingCostTypeRepositoryImplementation(db *gorm.DB) BookingCostTypeRepository {
	return &BookingCostTypeRepositoryImplementation{Db: db}
}

func (t *BookingCostTypeRepositoryImplementation) FindAll() []models.BookingCostType {
	var bookingCostTypes []models.BookingCostType
	result := t.Db.Find(&bookingCostTypes)
	if result.Error != nil {
		return []models.BookingCostType{}
	}

	return bookingCostTypes
}

func (t *BookingCostTypeRepositoryImplementation) FindById(id uint) models.BookingCostType {
	var bookingCostType models.BookingCostType
	result := t.Db.Where("id = ?", id).First(&bookingCostType)
	if result.Error != nil {
		return models.BookingCostType{}
	}

	return bookingCostType
}

func (t *BookingCostTypeRepositoryImplementation) Create(bookingCostType models.BookingCostType) models.BookingCostType {
	result := t.Db.Create(&bookingCostType)
	if result.Error != nil {
		return models.BookingCostType{}
	}

	return bookingCostType
}
