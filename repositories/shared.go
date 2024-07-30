package repositories

import (
	"booking-api/config"
	"booking-api/models"
	"fmt"

	"gorm.io/gorm"
)

func GetThumbnailForEntity(entityID uint, entityType string, db *gorm.DB) string {
	var entityPhoto models.EntityPhoto
	result := db.Model(&models.EntityPhoto{}).Where("entity_id = ? AND entity_type = ?", entityID, entityType).Preload("Photo").Limit(1).First(&entityPhoto)
	if result.Error != nil {
		return ""
	}

	// load config
	env, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	base := env.OBJECT_STORAGE_URL

	if entityPhoto.Photo.URL != "" {
		entityPhoto.Photo.URL = "https://" + base + "/" + entityPhoto.Photo.URL
	}

	return entityPhoto.Photo.URL
}
