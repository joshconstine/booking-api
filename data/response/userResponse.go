package response

type UserResponse struct {
	UserID             string            `json:"userId"`
	Username           string            `json:"username"`
	FirstName          string            `json:"firstName"`
	LastName           string            `json:"lastName"`
	Email              string            `json:"email"`
	PhoneNumber        string            `json:"phoneNumber"`
	ProfilePicture     string            `json:"profilePicture"`
	PreferredName      string            `json:"preferredName"`
	Bookings           []BookingResponse `json:"bookings"`
	Chats              []ChatResponse
	PermissionRequests []EntityBookingPermissionResponse
}
