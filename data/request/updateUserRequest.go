package request

import "github.com/google/uuid"

type UpdateUserRequest struct {
	Username string
	UserID   uuid.UUID
}
