package request

import "github.com/google/uuid"

type CreateUserRequest struct {
	UserID      uuid.UUID
	Email       string
	Username    string
	FirstName   string
	LastName    string
	PhoneNumber string
}
