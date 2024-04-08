package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	URL string
}

func (p *Photo) TableName() string {
	return "photos"
}

func (p *Photo) MapPhotoToResponse() response.PhotoResponse {
	return response.PhotoResponse{
		ID:  p.ID,
		URL: p.URL,
	}
}
