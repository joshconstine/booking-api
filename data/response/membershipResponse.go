package response

type MembershipResponse struct {
	ID          uint             `json:"id"`
	User        UserResponse     `json:"user"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Role        UserRoleResponse `json:"role"`
}
