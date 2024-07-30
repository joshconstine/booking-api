package response

type MembershipResponse struct {
	ID          uint             `json:"id"`
	UserID      string           `json:"userId"`
	AccountID   uint             `json:"accountId"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Role        UserRoleResponse `json:"role"`
}
