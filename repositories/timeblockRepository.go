package repositories

import (
	"booking-api/models"
)

type TimeblockRepository interface {
	FindAll() []models.EntityTimeblock
	FindByEntity(entityType string, entityId uint) []models.EntityTimeblock
	Create(timeblock models.EntityTimeblock) models.EntityTimeblock
}
