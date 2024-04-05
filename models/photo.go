package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	URL string
}

func (p *Photo) TableName() string {
	return "photos"
}
