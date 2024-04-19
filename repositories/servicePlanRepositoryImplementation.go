package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type ServicePlanRepositoryImplementation struct {
	DB *gorm.DB
}

func NewServicePlanRepositoryImplementation(db *gorm.DB) ServicePlanRepository {
	return &ServicePlanRepositoryImplementation{DB: db}
}

func (r *ServicePlanRepositoryImplementation) FindAll() []response.ServicePlanResponse {
	var servicePlans []models.ServicePlan
	var servicePlanResponses []response.ServicePlanResponse

	result := r.DB.Find(&servicePlans)
	if result.Error != nil {
		return []response.ServicePlanResponse{}
	}

	for _, servicePlan := range servicePlans {
		servicePlanResponses = append(servicePlanResponses, servicePlan.MapServicePlanToResponse())
	}

	return servicePlanResponses
}

func (r *ServicePlanRepositoryImplementation) Create(servicePlan models.ServicePlan) response.ServicePlanResponse {
	result := r.DB.Create(&servicePlan)
	if result.Error != nil {
		return response.ServicePlanResponse{}
	}

	return servicePlan.MapServicePlanToResponse()
}
