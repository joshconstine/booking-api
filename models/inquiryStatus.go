package models

import (
	"gorm.io/gorm"
)

type InquiryStatus struct {
	gorm.Model
	Name string `gorm:"unique; not null"`
}
