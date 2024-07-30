package response

type PublicUserResponse struct {
	Username       string `json:"username"`
	ProfilePicture string `json:"profilePicture"`
	PreferredName  string `json:"preferredName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
}
