package models

import "booking-api/data/response"

const UserContextKey = "user"

type AuthenticatedUser struct {
	User        response.UserResponse
	LoggedIn    bool
	AccessToken string
}
