package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
)

type UserRepository interface {
	FindAll() []models.User
	FindById(id uint) models.User
	FindByUserID(userID string) response.UserResponse
	IsAdmin(userID string) bool
	FindByEmail(email string) models.User
	Create(user *request.CreateUserRequest) error
	Update(user *request.UpdateUserRequest) error
}
