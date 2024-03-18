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

func (t *TimeblockRepositoryImplementation) FindAll() []models.Timeblock {
	var timeblocks []models.Timeblock
	result := t.Db.Find(&timeblocks)
	if result.Error != nil {
		return []models.Timeblock{}
	}

	return timeblocks
}

func (t *TimeblockRepositoryImplementation) FindByEntity(entityType string, entityId uint) []models.Timeblock {
	var timeblocks []models.Timeblock
	result := t.Db.Where("entity_type = ? AND entity_id = ?", entityType, entityId).Find(&timeblocks)
	if result.Error != nil {
		return []models.Timeblock{}
	}

	return timeblocks
}

func (t *TimeblockRepositoryImplementation) Create(timeblock models.Timeblock) models.Timeblock {
	result := t.Db.Create(&timeblock)
	if result.Error != nil {
		return models.Timeblock{}
	}

	return timeblock
}
