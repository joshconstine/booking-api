package repositories

import (
	"booking-api/data/request"
	"booking-api/models"

	"gorm.io/gorm"
)

type EntityTimeblockRepositoryImplementation struct {
	Db *gorm.DB
}

func NewTimeblockRepositoryImplementation(db *gorm.DB) EntityTimeblockRepository {
	return &EntityTimeblockRepositoryImplementation{Db: db}
}

func (t *EntityTimeblockRepositoryImplementation) FindAll() []models.EntityTimeblock {
	var timeblocks []models.EntityTimeblock
	result := t.Db.Find(&timeblocks)
	if result.Error != nil {
		return []models.EntityTimeblock{}
	}

	return timeblocks
}

func (t *EntityTimeblockRepositoryImplementation) FindByEntity(entityType string, entityId uint) []models.EntityTimeblock {
	var timeblocks []models.EntityTimeblock
	result := t.Db.Where("entity_type = ? AND entity_id = ?", entityType, entityId).Find(&timeblocks)
	if result.Error != nil {
		return []models.EntityTimeblock{}
	}

	return timeblocks
}

func (t *EntityTimeblockRepositoryImplementation) Create(timeblock request.CreateEntityTimeblockRequest) models.EntityTimeblock {
	timeblockToCreate := timeblock.MapCreateEntityTimeblockRequestToEntityTimeblock()
	result := t.Db.Create(&timeblockToCreate)
	if result.Error != nil {
		return models.EntityTimeblock{}
	}

	return timeblockToCreate
}
