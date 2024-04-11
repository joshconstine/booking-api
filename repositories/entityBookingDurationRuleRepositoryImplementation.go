package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type EntityBookingDurationRuleRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityBookingDurationRuleRepositoryImplementation(db *gorm.DB) EntityBookingDurationRuleRepository {
	return &EntityBookingDurationRuleRepositoryImplementation{Db: db}
}

func (e *EntityBookingDurationRuleRepositoryImplementation) FindById(entity_id uint, entity_type string) response.EntityBookingDurationRuleResponse {
	var entityBookingDurationRule models.EntityBookingDurationRule
	result := e.Db.Where("entity_id = ? AND entity_type = ?", entity_id, entity_type).First(&entityBookingDurationRule)
	if result.Error != nil {
		return response.EntityBookingDurationRuleResponse{}
	}

	return entityBookingDurationRule.MapEntityBookingDurationRuleToResponse()
}

func (e *EntityBookingDurationRuleRepositoryImplementation) Update(entityBookingDurationRule requests.UpdateEntityBookingDurationRuleRequest) (response.EntityBookingDurationRuleResponse, error) {
	var entityBookingDurationRuleToInsert models.EntityBookingDurationRule
	result := e.Db.Where("entity_id = ? AND entity_type = ?", entityBookingDurationRule.EntityID, entityBookingDurationRule.EntityType).First(&entityBookingDurationRuleToInsert)
	if result.Error != nil {
		return response.EntityBookingDurationRuleResponse{}, result.Error
	}

	entityBookingDurationRuleToInsert.EntityID = entityBookingDurationRule.EntityID
	entityBookingDurationRuleToInsert.EntityType = entityBookingDurationRule.EntityType
	entityBookingDurationRuleToInsert.MinimumDuration = entityBookingDurationRule.MinDuration
	entityBookingDurationRuleToInsert.MaximumDuration = entityBookingDurationRule.MaxDuration
	entityBookingDurationRuleToInsert.StartTime = entityBookingDurationRule.StartTime
	entityBookingDurationRuleToInsert.EndTime = entityBookingDurationRule.EndTime

	result = e.Db.Save(&entityBookingDurationRuleToInsert)
	if result.Error != nil {
		return response.EntityBookingDurationRuleResponse{}, result.Error
	}

	return entityBookingDurationRuleToInsert.MapEntityBookingDurationRuleToResponse(), nil

}
