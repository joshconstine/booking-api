package repositories

import (
	"booking-api/data/response"
)

type UserRoleRepository interface {
	FindAll() []response.UserRoleResponse
	FindByID(id uint) response.UserRoleResponse
	Create(name string) response.UserRoleResponse
}
