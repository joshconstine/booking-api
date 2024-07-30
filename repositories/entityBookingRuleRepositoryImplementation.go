package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type EntityBookingRuleRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityBookingRuleRepositoryImplementation(db *gorm.DB) EntityBookingRuleRepository {
	return &EntityBookingRuleRepositoryImplementation{Db: db}
}

func (e *EntityBookingRuleRepositoryImplementation) FindByID(entityID uint, entityType string) (response.EntityBookingRuleResponse, error) {
	var entityBookingRule models.EntityBookingRule
	result := e.Db.Where("entity_id = ? AND entity_type = ?", entityID, entityType).First(&entityBookingRule)
	if result.Error != nil {
		return response.EntityBookingRuleResponse{}, result.Error
	}

	return entityBookingRule.MapEntityBookingRuleToResponse(), nil
}

func (e *EntityBookingRuleRepositoryImplementation) Update(entityBookingRule request.UpdateEntityBookingRuleRequest) (response.EntityBookingRuleResponse, error) {
	var entityBookingRuleToInsert models.EntityBookingRule
	result := e.Db.Where("entity_id = ? AND entity_type = ?", entityBookingRule.EntityID, entityBookingRule.EntityType).First(&entityBookingRuleToInsert)
	if result.Error != nil {
		return response.EntityBookingRuleResponse{}, result.Error
	}

	entityBookingRuleToInsert.EntityID = entityBookingRule.EntityID
	entityBookingRuleToInsert.EntityType = entityBookingRule.EntityType
	entityBookingRuleToInsert.AdvertiseAtAllLocations = entityBookingRule.AdvertiseAtAllLocations
	entityBookingRuleToInsert.AllowPets = entityBookingRule.AllowPets
	entityBookingRuleToInsert.AllowInstantBooking = entityBookingRule.AllowInstantBooking
	entityBookingRuleToInsert.OfferEarlyCheckIn = entityBookingRule.OfferEarlyCheckIn

	result = e.Db.Save(&entityBookingRuleToInsert)
	if result.Error != nil {
		return response.EntityBookingRuleResponse{}, result.Error
	}

	return entityBookingRuleToInsert.MapEntityBookingRuleToResponse(), nil

}
