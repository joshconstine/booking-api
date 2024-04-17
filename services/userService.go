package services

import (
	requests "booking-api/data/request"
	responses "booking-api/data/response"

	"github.com/google/uuid"
)

type UserService interface {
	FindAll() []responses.UserResponse
	FindByEmail(email string) responses.UserResponse
	FindByUserID(userID uuid.UUID) responses.UserResponse
	CreateUser(request *requests.CreateUserRequest) error
	UpdateUser(user *requests.UpdateUserRequest) error
}
