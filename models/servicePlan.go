package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type ServicePlan struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
	Fees []ServiceFee
}

func (s *ServicePlan) TableName() string {

	return "service_plans"
}

func (s *ServicePlan) MapServicePlanToResponse() response.ServicePlanResponse {
	var fees []response.ServiceFeeResponse
	for _, fee := range s.Fees {
		fees = append(fees, fee.MapServiceFeeToResponse())
	}
	return response.ServicePlanResponse{
		ID:   s.ID,
		Name: s.Name,
		Fees: fees,
	}
}
