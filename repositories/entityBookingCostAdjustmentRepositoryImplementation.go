package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type EntityBookingCostAdjustmentRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityBookingCostAdjustmentRepositoryImplementation(db *gorm.DB) EntityBookingCostAdjustmentRepository {
	return &EntityBookingCostAdjustmentRepositoryImplementation{Db: db}
}

func (t *EntityBookingCostAdjustmentRepositoryImplementation) FindAllForEntity(entityID uint, entityType string) []response.EntityBookingCostAdjustmentResponse {
	var entityBookingCostAdjustments []models.EntityBookingCostAdjustment
	t.Db.Where("entity_id = ? AND entity_type = ?", entityID, entityType).Find(&entityBookingCostAdjustments)

	var entityBookingCostAdjustmentResponses []response.EntityBookingCostAdjustmentResponse
	for _, entityBookingCostAdjustment := range entityBookingCostAdjustments {
		entityBookingCostAdjustmentResponses = append(entityBookingCostAdjustmentResponses, entityBookingCostAdjustment.MapEntityBookingCostAdjustmentToResponse())
	}

	return entityBookingCostAdjustmentResponses
}

func (t *EntityBookingCostAdjustmentRepositoryImplementation) FindAllForEntityAndRange(entityID uint, entityType string, startDate string, endDate string) []response.EntityBookingCostAdjustmentResponse {
	var entityBookingCostAdjustments []models.EntityBookingCostAdjustment
	t.Db.Where("entity_id = ? AND entity_type = ? AND start_date >= ? AND end_date <= ?", entityID, entityType, startDate, endDate).Find(&entityBookingCostAdjustments)

	var entityBookingCostAdjustmentResponses []response.EntityBookingCostAdjustmentResponse
	for _, entityBookingCostAdjustment := range entityBookingCostAdjustments {
		entityBookingCostAdjustmentResponses = append(entityBookingCostAdjustmentResponses, entityBookingCostAdjustment.MapEntityBookingCostAdjustmentToResponse())
	}

	return entityBookingCostAdjustmentResponses
}

func (t *EntityBookingCostAdjustmentRepositoryImplementation) Create(request request.CreateEntityBookingCostAdjustmentRequest) (response.EntityBookingCostAdjustmentResponse, error) {
	entityBookingCostAdjustment := models.EntityBookingCostAdjustment{
		EntityID:          request.EntityID,
		EntityType:        request.EntityType,
		BookingCostTypeID: request.BookingCostTypeID,
		Amount:            request.Amount,
		StartDate:         request.StartDate,
		EndDate:           request.EndDate,
		TaxRateID:         request.TaxRateID,
	}
	result := t.Db.Create(&entityBookingCostAdjustment)
	if result.Error != nil {
		return response.EntityBookingCostAdjustmentResponse{}, result.Error
	}

	return entityBookingCostAdjustment.MapEntityBookingCostAdjustmentToResponse(), nil
}

func (t *EntityBookingCostAdjustmentRepositoryImplementation) Update(request request.UpdateEntityBookingCostAdjustmentRequest) (response.EntityBookingCostAdjustmentResponse, error) {

	var entityBookingCostAdjustmentToUpdate models.EntityBookingCostAdjustment
	result := t.Db.Model(&models.EntityBookingCostAdjustment{}).Where("id = ?", request.ID).First(&entityBookingCostAdjustmentToUpdate)

	if result.Error != nil {
		return response.EntityBookingCostAdjustmentResponse{}, result.Error
	}

	if entityBookingCostAdjustmentToUpdate == (models.EntityBookingCostAdjustment{}) {
		// return response.EntityBookingCostAdjustmentResponse{}, error.Error("Entity Booking Cost Adjustment not found")
		return response.EntityBookingCostAdjustmentResponse{}, nil
	}

	entityBookingCostAdjustmentToUpdate.Amount = request.Amount
	entityBookingCostAdjustmentToUpdate.StartDate = request.StartDate
	entityBookingCostAdjustmentToUpdate.EndDate = request.EndDate
	entityBookingCostAdjustmentToUpdate.TaxRateID = request.TaxRateID

	result = t.Db.Save(&entityBookingCostAdjustmentToUpdate)
	if result.Error != nil {
		return response.EntityBookingCostAdjustmentResponse{}, result.Error
	}

	return entityBookingCostAdjustmentToUpdate.MapEntityBookingCostAdjustmentToResponse(), nil
}

func (t *EntityBookingCostAdjustmentRepositoryImplementation) Delete(adjustmentId uint) error {
	result := t.Db.Model(&models.EntityBookingCostAdjustment{}).Where("id = ?", adjustmentId).Delete(&models.EntityBookingCostAdjustment{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
