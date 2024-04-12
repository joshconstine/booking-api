package models

const UserContextKey = "user"

type AuthenticatedUser struct {
	User
	LoggedIn bool
}
