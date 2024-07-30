package repositories

import (
	"booking-api/data/request"
	"booking-api/models"
)

type EntityTimeblockRepository interface {
	FindAll() []models.EntityTimeblock
	FindByEntity(entityType string, entityId uint) []models.EntityTimeblock
	Create(timeblock request.CreateEntityTimeblockRequest) models.EntityTimeblock
}
