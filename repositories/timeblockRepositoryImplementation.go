package repositories

import (
	"booking-api/models"

	"gorm.io/gorm"
)

type TimeblockRepositoryImplementation struct {
	Db *gorm.DB
}

func NewTimeblockRepositoryImplementation(db *gorm.DB) TimeblockRepository {
	return &TimeblockRepositoryImplementation{Db: db}
}

func (t *TimeblockRepositoryImplementation) FindAll() []models.EntityTimeblock {
	var timeblocks []models.EntityTimeblock
	result := t.Db.Find(&timeblocks)
	if result.Error != nil {
		return []models.EntityTimeblock{}
	}

	return timeblocks
}

func (t *TimeblockRepositoryImplementation) FindByEntity(entityType string, entityId uint) []models.EntityTimeblock {
	var timeblocks []models.EntityTimeblock
	result := t.Db.Where("entity_type = ? AND entity_id = ?", entityType, entityId).Find(&timeblocks)
	if result.Error != nil {
		return []models.EntityTimeblock{}
	}

	return timeblocks
}

func (t *TimeblockRepositoryImplementation) Create(timeblock models.EntityTimeblock) models.EntityTimeblock {
	result := t.Db.Create(&timeblock)
	if result.Error != nil {
		return models.EntityTimeblock{}
	}

	return timeblock
}
