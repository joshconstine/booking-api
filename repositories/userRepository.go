package repositories

import (
	"booking-api/data/request"
	"booking-api/models"
)

type UserRepository interface {
	FindAll() []models.User
	FindById(id uint) models.User
	FindByEmail(email string) models.User
	Create(user request.CreateUserRequest) int
	Update(user models.User) models.User
}
