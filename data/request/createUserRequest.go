package request

type CreateUserRequest struct {
	UserID      string
	Email       string
	Username    string
	FirstName   string
	LastName    string
	PhoneNumber string
}
