package repositories

import (
	"booking-api/models"
)

type BookingCostTypeRepository interface {
	FindAll() []models.BookingCostType
	FindById(id uint) models.BookingCostType
	Create(bookingCostType models.BookingCostType) models.BookingCostType
}
