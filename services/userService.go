package services

import (
	requests "booking-api/data/request"
	responses "booking-api/data/response"
)

type UserService interface {
	FindByEmail(email string) responses.UserResponse
	CreateUser(request requests.CreateUserRequest) responses.UserResponse
}
