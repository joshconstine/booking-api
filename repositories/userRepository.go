package repositories

import (
	"booking-api/data/request"
	"booking-api/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() []models.User
	FindById(id uint) models.User
	FindByUserID(userID uuid.UUID) models.User
	FindByEmail(email string) models.User
	Create(user *request.CreateUserRequest) error
	Update(user models.User) models.User
}
