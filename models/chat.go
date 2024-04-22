package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	AccountID uint `gorm:"not null"`
	UserID    uint `gorm:"not null"`
	Messages  []ChatMessage
}
