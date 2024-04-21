package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"fmt"
	"log/slog"
	"strconv"

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
		"id = ?", id).
		Preload("Payments.PaymentMethod").
		Preload("Details").
		Preload("CostItems").
		Preload("User").
		Preload("CostItems.BookingCostType").
		Preload("CostItems.TaxRate").
		Preload("Documents.Document").
		Preload("BookingStatus").
		Preload("Entities").
		First(&booking)

	if result.Error != nil {
		return response.BookingInformationResponse{}
	}

	return booking.MapBookingToInformationResponse()

}

func (t *bookingRepositoryImplementation) Create(booking *models.Booking) error {

	result := t.Db.Model(&models.Booking{}).Create(&booking)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return result.Error
	}

	return nil
}

func (t *bookingRepositoryImplementation) Update(booking models.Booking) models.Booking {
	result := t.Db.Save(&booking)
	if result.Error != nil {
		return models.Booking{}
	}

	return booking
}
func (t *bookingRepositoryImplementation) CheckIfEntitiesCanBeBooked(request *request.CreateBookingRequest) (bool, error) {
	for _, entity := range request.EntityRequests {
		var entityTimeblocks []models.EntityTimeblock
		result := t.Db.Model(&models.EntityTimeblock{}).Where(
			"entity_id = ? AND entity_type = ?",
			entity.EntityID,
			entity.EntityType,
		).Find(&entityTimeblocks)

		if result.Error != nil {
			return false, result.Error
		}

		for _, timeblock := range entityTimeblocks {
			if (entity.StartTime.Equal(timeblock.StartTime) || entity.StartTime.After(timeblock.StartTime)) && (entity.StartTime.Before(timeblock.EndTime)) ||
				((entity.EndTime.Equal(timeblock.EndTime) || entity.EndTime.Before(timeblock.EndTime)) && (entity.EndTime.After(timeblock.StartTime))) ||
				(entity.StartTime.Before(timeblock.StartTime) && entity.EndTime.After(timeblock.EndTime)) {
				errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) is already booked for the time range %s to %s", strconv.Itoa(int(entity.EntityID)), entity.EntityType, entity.StartTime, entity.EndTime)
				return false, fmt.Errorf(errorMsg)
			}
		}
	}

	return true, nil
}
