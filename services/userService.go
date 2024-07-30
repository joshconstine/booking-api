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
	FindByPublicUserID(publicUserID string) (responses.UserResponse, error)
	IsAdmin(userID string) bool
	IsOwnerOfEntity(userID string, entityType string, entityID uint) (bool, error)
	CreateUser(request *requests.CreateUserRequest) error
	UpdateUser(user *requests.UpdateUserRequest) error
}
