package request

import "github.com/google/uuid"

type UpdateUserRequest struct {
	UserID      uuid.UUID
	Username    string
	FirstName   string
	LastName    string
	PhoneNumber string
}
