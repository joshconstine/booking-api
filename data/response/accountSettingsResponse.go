package response

type AccountSettingsResponse struct {
	ID           uint `json:"id"`
	AccountID    uint `json:"accountID"`
	ServicePlan  ServicePlanResponse
	AccountOwner MembershipResponse
}
