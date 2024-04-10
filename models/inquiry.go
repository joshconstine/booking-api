package models

import (
	"gorm.io/gorm"
)

type Inquiry struct {
	gorm.Model
	UserID          uint `gorm:"not null"`
	Note            string
	NumGuests       int
	User            User
	EntityInquiries []EntityInquiry
}
