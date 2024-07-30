package response

type AccountResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Members         []MembershipResponse
	AccountSettings AccountSettingsResponse
}
