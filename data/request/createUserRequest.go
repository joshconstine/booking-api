package request

import "github.com/google/uuid"

type CreateUserRequest struct {
	UserID   uuid.UUID `json:"userId"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
