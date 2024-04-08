package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type RentalStatus struct {
	gorm.Model
	RentalID uint `gorm:"uniqueIndex"`
	IsClean  bool
}

func (r *RentalStatus) TableName() string {
	return "rental_statuses"
}

func (r *RentalStatus) MapRentalStatusToResponse() response.RentalStatusResponse {
	return response.RentalStatusResponse{
		IsClean: r.IsClean,
	}
}
