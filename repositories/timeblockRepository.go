package repositories

import (
	"booking-api/models"
)

type TimeblockRepository interface {
	FindAll() []models.Timeblock
	FindByEntity(entityType string, entityId uint) []models.Timeblock
	Create(timeblock models.Timeblock) models.Timeblock
}
