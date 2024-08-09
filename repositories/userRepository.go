package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
)

type UserRepository interface {
	FindAll() []models.User
	FindById(id uint) (response.UserResponse, error)
	FindByPublicUserID(publicUserID string) (response.UserResponse, error)
	FindByUserID(userID string) (response.UserResponse, error)
	IsAdmin(userID string) bool
	FindByEmail(email string) (models.User, error)
	Create(user *request.CreateUserRequest) error
	CreateForUser(user *request.CreateUserRequestForUser) error
	Update(user *request.UpdateUserRequest) error
}
