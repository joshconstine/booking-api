package services

import (
	"booking-api/data/response"
)

type UserRoleService interface {
	FindAll() []response.UserRoleResponse
	FindByID(id uint) response.UserRoleResponse
	Create(name string) response.UserRoleResponse
}
