package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type ServiceFee struct {
	gorm.Model
	ServicePlanID         uint    `gorm:"not null"`
	FeePercentage         float64 `gorm:"not null"`
	AppliesToAllCostTypes bool
	BookingCostTypeID     uint
}

func (s *ServiceFee) MapServiceFeeToResponse() response.ServiceFeeResponse {
	return response.ServiceFeeResponse{
		ID:                    s.ID,
		FeePercentage:         s.FeePercentage,
		AppliesToAllCostTypes: s.AppliesToAllCostTypes,
		BookingCostTypeID:     s.BookingCostTypeID,
	}
}
