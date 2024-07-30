package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type UserRoleRepository interface {
	FindAll() []response.UserRoleResponse
	FindByID(id uint) response.UserRoleResponse
	Create(userRole *models.UserRole) response.UserRoleResponse
}
