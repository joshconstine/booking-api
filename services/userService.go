package services

import (
	requests "booking-api/data/request"
	responses "booking-api/data/response"
)

type UserService interface {
	FindAll() []responses.UserResponse
	FindByEmail(email string) responses.UserResponse
	FindByUserID(userID string) responses.UserResponse
	CreateUser(request *requests.CreateUserRequest) error
	UpdateUser(user *requests.UpdateUserRequest) error
}
