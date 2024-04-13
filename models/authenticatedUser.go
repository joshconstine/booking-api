package models

const UserContextKey = "user"

type AuthenticatedUser struct {
	User        User
	LoggedIn    bool
	AccessToken string
}
