package models

import (
	"booking-api/config"
	"booking-api/data/response"
	"path"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	URL string `gorm:"not null"`
}

func (p *Photo) TableName() string {
	return "photos"
}

func (p *Photo) MapPhotoToResponse() response.PhotoResponse {

	// load config
	env, err := config.LoadConfig(".")
	if err != nil {
		return response.PhotoResponse{}
	}
	urlstring := path.Join("https://", env.OBJECT_STORAGE_URL, p.URL)
	return response.PhotoResponse{
		ID:  p.ID,
		URL: urlstring,
	}
}
