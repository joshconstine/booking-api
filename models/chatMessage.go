package models

import (
	"gorm.io/gorm"
)

type ChatMessage struct {
	gorm.Model
	ChatID  uint   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
	Message string `gorm:"not null"`
}
