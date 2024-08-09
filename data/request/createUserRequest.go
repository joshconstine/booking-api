package request

import "time"

type CreateUserRequest struct {
	UserID         string
	Email          string
	Username       string
	FirstName      string
	LastName       string
	PhoneNumber    string
	DOB            *time.Time
	ProfilePicture string
}
type CreateUserRequestForUser struct {
	Email     string
	FirstName string
	LastName  string
}
