package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Members  []Membership
	Settings AccountSettings
}
