package repositories

import (
	"booking-api/constants"
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
		BookingID:       entityBooking.BookingID,
		EntityID:        entityBooking.EntityID,
		EntityType:      entityBooking.EntityType,
		BookingStatusID: constants.BOOKING_STATUS_DRAFTED_ID,
		Timeblock: models.EntityTimeblock{
			StartTime:  entityBooking.StartTime,
			EndTime:    entityBooking.EndTime,
			EntityID:   entityBooking.EntityID,
			EntityType: entityBooking.EntityType,
		},
		BookingCostItems: []models.BookingCostItem{},
		Documents:        []models.BookingDocument{},
	}

	result := e.Db.Create(&entityBookingModel)
	if result.Error != nil {
		return response.EntityBookingResponse{}
	}
	//Get the cost items for the entity
	var entityCostItems []models.EntityBookingCost

	result = e.Db.Model(&models.EntityBookingCost{}).Where("entity_id = ? AND entity_type = ?", entityBooking.EntityID, entityBooking.EntityType).Find(&entityCostItems)
	if result.Error != nil {
		return response.EntityBookingResponse{}
	}

	//Get The documents for the entity

	var entityDocuments []models.EntityBookingDocument
	result = e.Db.Model(&models.EntityBookingDocument{}).Where("entity_id = ? AND entity_type = ?", entityBooking.EntityID, entityBooking.EntityType).Find(&entityDocuments)
	if result.Error != nil {
		return response.EntityBookingResponse{}
	}

	// var bookingDocuments []models.BookingDocument
	// var bookingDocument models.BookingDocument
	// for _, document := range entityDocuments {
	// 	bookingDocument = document.MapEntityBookingDocumentToBookingDocument(entityBookingModel.BookingID, entityBookingModel.ID)
	// 	bookingDocuments = append(bookingDocuments, bookingDocument)

	// }

	var bookingCostItems []models.BookingCostItem
	var bookingCostItem models.BookingCostItem
	for _, costItem := range entityCostItems {
		bookingCostItem = costItem.MapEntityBookingCostToBookingCostItem(entityBooking.BookingID, entityBookingModel.ID)
		bookingCostItems = append(bookingCostItems, bookingCostItem)
	}

	entityBookingModel.BookingCostItems = bookingCostItems
	// entityBookingModel.Documents = bookingDocuments

	result = e.Db.Save(&entityBookingModel)
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
func (e *EntityBookingRepositoryImplementation) UpdateStatus(request request.UpdateEntityBookingStatusRequest) error {
	var entityBooking models.EntityBooking
	result := e.Db.Model(&models.EntityBooking{}).Where("id = ?", request.EntityBookingID).First(&entityBooking)
	if result.Error != nil {
		return result.Error
	}
	entityBooking.BookingStatusID = request.BookingStatusID
	result = e.Db.Save(&entityBooking)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
