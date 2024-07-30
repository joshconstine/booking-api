package models

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	Iso       string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null"`
}

func (c *Country) TableName() string {
	return "countries"
}
