package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() []models.User
	FindById(id uint) models.User
	FindByUserID(userID uuid.UUID) response.UserResponse
	FindByEmail(email string) models.User
	Create(user *request.CreateUserRequest) error
	Update(user *request.UpdateUserRequest) error
}
