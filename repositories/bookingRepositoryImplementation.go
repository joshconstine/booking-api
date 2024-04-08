package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type bookingRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingRepositoryImplementation(Db *gorm.DB) BookingRepository {
	return &bookingRepositoryImplementation{Db: Db}
}

func (t *bookingRepositoryImplementation) FindAll() []response.BookingResponse {
	var bookings []models.Booking
	result := t.Db.Find(&bookings)
	if result.Error != nil {
		return []response.BookingResponse{}
	}

	var response []response.BookingResponse
	for _, booking := range bookings {
		response = append(response, booking.MapBookingToResponse())
	}

	return response

}

func (t *bookingRepositoryImplementation) FindById(id string) response.BookingInformationResponse {
	var booking models.Booking
	result := t.Db.Model(&models.Booking{}).Where(
		"id = ?", id).Preload("Payments.PaymentMethod").Preload("Details").Preload("CostItems").Preload("Documents").Preload("BookingStatus").First(&booking)

	if result.Error != nil {
		return response.BookingInformationResponse{}
	}

	return booking.MapBookingToInformationResponse()

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
