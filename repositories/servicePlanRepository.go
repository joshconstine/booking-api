package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type ServicePlanRepository interface {
	FindAll() []response.ServicePlanResponse
	Create(servicePlan models.ServicePlan) response.ServicePlanResponse
}
