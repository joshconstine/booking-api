package services

import (
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/repositories"
)

type UserRoleServiceImplementation struct {
	userRoleRepository repositories.UserRoleRepository
}

func NewUserRoleServiceImplementation(userRoleRepository repositories.UserRoleRepository) UserRoleService {
	return &UserRoleServiceImplementation{
		userRoleRepository: userRoleRepository,
	}
}

func (r *UserRoleServiceImplementation) FindAll() []response.UserRoleResponse {
	return r.userRoleRepository.FindAll()
}

func (r *UserRoleServiceImplementation) FindByID(id uint) response.UserRoleResponse {
	return r.userRoleRepository.FindByID(id)
}

func (r *UserRoleServiceImplementation) Create(name string) response.UserRoleResponse {
	role := models.UserRole{
		Name: name,
	}
	return r.userRoleRepository.Create(&role)
}
