package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
)

type UserRepository interface {
	FindAll() []models.User
	FindById(id uint) models.User
	FindByPublicUserID(publicUserID string) (response.UserResponse, error)
	FindByUserID(userID string) response.UserResponse
	IsAdmin(userID string) bool
	FindByEmail(email string) (models.User, error)
	Create(user *request.CreateUserRequest) error
	Update(user *request.UpdateUserRequest) error
}
