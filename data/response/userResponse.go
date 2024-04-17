package response

import "github.com/google/uuid"

type UserResponse struct {
	UserID   uuid.UUID `json:"userId"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}
