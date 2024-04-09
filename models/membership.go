package models

import "gorm.io/gorm"

type Membership struct {
	gorm.Model
	AccountID   uint `gorm:"not null"`
	UserID      uint `gorm:"not null"`
	PhoneNumber string
	Email       string
	RoleID      uint `gorm:"not null"`
	Role        UserRole
	User        User
}
