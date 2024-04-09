package response

type ServicePlanResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Fees []ServiceFeeResponse
}
