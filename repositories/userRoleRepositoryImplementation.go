package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type UserRoleRepositoryImplementation struct {
	DB *gorm.DB
}

func NewUserRoleRepositoryImplementation(db *gorm.DB) UserRoleRepository {
	return &UserRoleRepositoryImplementation{
		DB: db,
	}
}

func (u *UserRoleRepositoryImplementation) FindAll() []response.UserRoleResponse {
	var userRoles []models.UserRole
	u.DB.Find(&userRoles)
	var userRoleResponses []response.UserRoleResponse
	for _, userRole := range userRoles {
		userRoleResponses = append(userRoleResponses, userRole.MapUserRoleToResponse())
	}
	return userRoleResponses
}

func (u *UserRoleRepositoryImplementation) FindByID(id uint) response.UserRoleResponse {
	var userRole models.UserRole
	u.DB.Where("id = ?", id).Find(&userRole)
	return userRole.MapUserRoleToResponse()
}

func (u *UserRoleRepositoryImplementation) Create(userRole *models.UserRole) response.UserRoleResponse {

	u.DB.Create(&userRole)
	return userRole.MapUserRoleToResponse()
}
