package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"time"

	"gorm.io/gorm"
)

type EntityBookingRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityBookingRepositoryImplementation(db *gorm.DB) EntityBookingRepository {
	return &EntityBookingRepositoryImplementation{Db: db}
}

func (e *EntityBookingRepositoryImplementation) FindAllForEntity(entityType string, entityID uint) []response.EntityBookingResponse {
	var entityBookings []response.EntityBookingResponse
	result := e.Db.Model(&models.EntityBooking{}).Where("entity_id = ? AND entity_type = ?", entityID, entityType).Find(&entityBookings)
	if result.Error != nil {
		return []response.EntityBookingResponse{}
	}

	return entityBookings
}

func (e *EntityBookingRepositoryImplementation) FindById(id uint) response.EntityBookingResponse {
	var entityBooking models.EntityBooking
	result := e.Db.Model(&models.EntityBooking{}).Where("id = ?", id).First(&entityBooking)
	if result.Error != nil {
		return response.EntityBookingResponse{}
	}

	return entityBooking.MapEntityBookingToResponse()
}

func (e *EntityBookingRepositoryImplementation) FindAllForBooking(bookingID string) []response.EntityBookingResponse {
	var entityBookings []response.EntityBookingResponse
	result := e.Db.Model(&models.EntityBooking{}).Where("booking_id = ?", bookingID).Find(&entityBookings)
	if result.Error != nil {
		return []response.EntityBookingResponse{}
	}

	return entityBookings
}

func (e *EntityBookingRepositoryImplementation) Create(entityBooking request.CreateEntityBookingRequest) response.EntityBookingResponse {
	entityBookingModel := models.EntityBooking{
		BookingID:  entityBooking.BookingID,
		EntityID:   entityBooking.EntityID,
		EntityType: entityBooking.EntityType,
		Timeblock: models.Timeblock{
			StartTime:  entityBooking.StartTime,
			EndTime:    entityBooking.EndTime,
			EntityID:   entityBooking.EntityID,
			EntityType: entityBooking.EntityType,
			BookingID:  entityBooking.BookingID,
		},
		BookingStatusID:  0,
		BookingCostItems: []models.BookingCostItem{},
	}
	result := e.Db.Create(&entityBookingModel)
	if result.Error != nil {
		return response.EntityBookingResponse{}
	}

	return entityBookingModel.MapEntityBookingToResponse()

}

func (e *EntityBookingRepositoryImplementation) Update(entityBooking request.UpdateEntityBookingRequest) response.EntityBookingResponse {
	entityBookingModel := models.EntityBooking{
		BookingID:  entityBooking.BookingID,
		EntityID:   entityBooking.EntityID,
		EntityType: entityBooking.EntityType,

		BookingStatusID:  entityBooking.BookingStatusID,
		BookingCostItems: []models.BookingCostItem{},
	}
	result := e.Db.Save(&entityBookingModel)
	if result.Error != nil {
		return response.EntityBookingResponse{}
	}

	return entityBookingModel.MapEntityBookingToResponse()
}

func (e *EntityBookingRepositoryImplementation) FindAllForEntityForRange(entityType string, entityID uint, startTime *time.Time, endTime *time.Time) []response.EntityBookingResponse {
	var entityBookings []models.EntityBooking
	var bookingsMatchingRange []response.EntityBookingResponse
	result := e.Db.Model(&models.EntityBooking{}).Where("entity_id = ? AND entity_type = ?", entityID, entityType).Find(&entityBookings)
	if result.Error != nil {
		return []response.EntityBookingResponse{}
	}

	for _, booking := range entityBookings {
		if booking.Timeblock.StartTime.After(*startTime) && booking.Timeblock.EndTime.Before(*endTime) {
			bookingsMatchingRange = append(bookingsMatchingRange, booking.MapEntityBookingToResponse())

		}

	}
	return bookingsMatchingRange

}
