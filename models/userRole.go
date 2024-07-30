package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (u *UserRole) TableName() string {
	return "user_roles"
}

func (u *UserRole) MapUserRoleToResponse() response.UserRoleResponse {
	return response.UserRoleResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
