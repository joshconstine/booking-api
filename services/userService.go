package services

import (
	requests "booking-api/data/request"
	responses "booking-api/data/response"
)

type UserService interface {
	FindAll() []responses.UserResponse
	FindByEmailPublic(email string) (responses.PublicUserResponse, error)
	FindByEmail(email string) (responses.UserResponse, error)
	FindByUserID(userID string) responses.UserResponse
	IsAdmin(userID string) bool
	CreateUser(request *requests.CreateUserRequest) error
	UpdateUser(user *requests.UpdateUserRequest) error
}
