package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"errors"

	"gorm.io/gorm"
)

type EntityBookingCostRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityBookingCostRepositoryImplementation(db *gorm.DB) EntityBookingCostRepository {
	return &EntityBookingCostRepositoryImplementation{Db: db}
}

func (e *EntityBookingCostRepositoryImplementation) FindAllForEntity(entityID uint, entityType string) []response.EntityBookingCostResponse {
	var entityBookingCosts []models.EntityBookingCost
	e.Db.Where("entity_id = ? AND entity_type = ?", entityID, entityType).
		Preload("BookingCostType").
		Preload("TaxRate").
		Find(&entityBookingCosts)

	var entityBookingCostResponses []response.EntityBookingCostResponse
	for _, entityBookingCost := range entityBookingCosts {
		entityBookingCostResponses = append(entityBookingCostResponses, entityBookingCost.MapEntityBookingCostToResponse())
	}

	return entityBookingCostResponses
}

func (e *EntityBookingCostRepositoryImplementation) Create(request request.CreateEntityBookingCostRequest) (response.EntityBookingCostResponse, error) {
	entityBookingCost := models.EntityBookingCost{
		EntityID:          request.EntityID,
		EntityType:        request.EntityType,
		BookingCostTypeID: request.BookingCostTypeID,
		Amount:            request.Amount,
		TaxRateID:         request.TaxRateID,
	}
	result := e.Db.Create(&entityBookingCost)
	if result.Error != nil {
		return response.EntityBookingCostResponse{}, result.Error
	}

	return entityBookingCost.MapEntityBookingCostToResponse(), nil
}

func (e *EntityBookingCostRepositoryImplementation) Update(request request.UpdateEntityBookingCostRequest) (response.EntityBookingCostResponse, error) {

	var entityBookingCostToUpdate models.EntityBookingCost
	result := e.Db.Where("entity_id = ? AND entity_type = ? AND booking_cost_type_id = ?", request.EntityID, request.EntityType, request.BookingCostTypeID).First(&entityBookingCostToUpdate)
	if result.Error != nil {
		return response.EntityBookingCostResponse{}, result.Error
	}

	if entityBookingCostToUpdate == (models.EntityBookingCost{}) {
		// return response.EntityBookingCostResponse{}, error.Error("Entity Booking Cost not found")
		return response.EntityBookingCostResponse{}, errors.New("Entity Booking Cost not found")
	}

	entityBookingCostToUpdate.Amount = request.Amount
	entityBookingCostToUpdate.TaxRateID = request.TaxRateID

	result = e.Db.Save(&entityBookingCostToUpdate)
	if result.Error != nil {
		return response.EntityBookingCostResponse{}, result.Error
	}

	return entityBookingCostToUpdate.MapEntityBookingCostToResponse(), nil
}

func (e *EntityBookingCostRepositoryImplementation) Delete(entityID uint, entityType string, bookingCostTypeID uint) error {
	result := e.Db.Where("entity_id = ? AND entity_type = ? AND booking_cost_type_id = ?", entityID, entityType, bookingCostTypeID).Delete(&models.EntityBookingCost{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
